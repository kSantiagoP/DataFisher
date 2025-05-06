package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/data_api"
	"github.com/kSantiagoP/DataFisher/internal/logger"
	companyEmail "github.com/kSantiagoP/DataFisher/internal/model/company_email"
	"github.com/kSantiagoP/DataFisher/internal/model/job"
)

func EnrichCnpj(job []byte) error {
	logg := logger.NewLogger("workerProcessor")

	request := JobRequest{}
	err := json.Unmarshal(job, &request)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	tracker := config.GetRedisTracker()

	maxRetries := 3
	for _, cnpj := range request.Cnpjs {
		//primeiro checar se esse cnpj ja foi enriquecido
		retryCount := 0
		var lastError error

		for retryCount < maxRetries {
			err := enrichCNPJ(cnpj)
			if err == nil { //em sucesso, prossiga, caso contrário tente até o num maximo de vezes
				err = tracker.IncrementProgress(request.JobId)
				if err != nil {
					logg.Errorf("error updating progress: %v", err)
				}

				break
			}

			lastError = err
			retryCount++
			logg.Infof("Retry %d for CNPJ %s (error: %v)", retryCount, cnpj, err)
			time.Sleep(time.Second * time.Duration(retryCount))
		}

		if retryCount == maxRetries {
			logg.Errorf("max retries reached for cnpj %s, %v", cnpj, err)
			err = tracker.MarkFailedCNPJ(request.JobId, cnpj, lastError)
			if err != nil {
				logg.Errorf("error marking failed cnpj: %v", err)
			}
		}
	}

	failed, err := tracker.GetFailedCount(request.JobId)
	if err != nil {
		return fmt.Errorf("error checking failed jobs: %v", err)
	}

	if failed > 0 {

		if err := recordJobCnpj(request.JobId, request.Cnpjs); err != nil {
			logg.Errorf("error saving completed job in database: %v", err)
			return err
		}

		err := tracker.PartiallyCompleteJob(request.JobId, failed)
		if err != nil {
			return fmt.Errorf("error marking partial completion: %v", err)
		}

		err = recordJobStatus(request.JobId, tracker)
		if err != nil {
			return fmt.Errorf("error saving job into database: %v", err)
		}

		logg.Infof("job %s completed with %d failures", request.JobId, failed)
		return nil
	}

	if err = recordJobCnpj(request.JobId, request.Cnpjs); err != nil {
		logg.Errorf("error saving completed job in database: %v", err)
		return err
	}

	err = recordJobStatus(request.JobId, tracker)
	if err != nil {
		return fmt.Errorf("error saving job into database: %v", err)
	}

	logg.Infof("job %s completed successfully", request.JobId)
	return tracker.CompleteJob(request.JobId)
}

func enrichCNPJ(cnpj string) error {
	db := config.GetPostgresDB()
	emailList := data_api.GetEmailsByCnpj(cnpj)
	if len(emailList) == 0 {
		return fmt.Errorf("error: cnpj doesn't have emails registered")
	}
	emailsEnriched := data_api.ValidateEmail(emailList, cnpj)

	if err := db.Create(&emailsEnriched).Error; err != nil {
		return fmt.Errorf("error creating company Email: %v", err.Error())
	}

	phoneList := data_api.GetPhonesByCnpj(cnpj)
	if len(phoneList) == 0 {
		return fmt.Errorf("error: cnpj doesn't have phones registered")
	}
	phonesEnriched := data_api.ValidatePhone(phoneList, cnpj)

	if err := db.Create(&phonesEnriched).Error; err != nil {
		db.Where("cnpj = ?", cnpj).Delete(&companyEmail.CompanyEmail{})
		return fmt.Errorf("phone creation failed: %v", err)
	}

	return nil
}

func recordJobCnpj(jobId string, cnpjs []string) error {
	jobBatch := []job.JobCnpj{}
	for _, cnpj := range cnpjs {
		jobBatch = append(jobBatch, job.JobCnpj{
			Cnpj:  cnpj,
			JobId: jobId,
		})
	}

	db := config.GetPostgresDB()

	return db.Create(&jobBatch).Error
}

func recordJobStatus(jobId string, tracker *config.JobTracker) error {
	data, err := tracker.GetJobStatus(jobId)
	if err != nil {
		return err
	}

	jobStatus := job.JobStatus{
		JobId:     data["job_id"].(string),
		Status:    data["status"].(string),
		Progress:  data["progress"].(float64),
		Companies: data["companies"].(int),
		Completed: data["completed"].(int),
		Failed:    data["failed"].(int),
		Pending:   data["pending"].(int),
	}

	db := config.GetPostgresDB()
	return db.Create(&jobStatus).Error
}
