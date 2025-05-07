package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/model/company"
	"github.com/kSantiagoP/DataFisher/internal/model/job"
	"github.com/kSantiagoP/DataFisher/internal/router/response"
	"github.com/kSantiagoP/DataFisher/internal/types"
)

func GetEnriquecimentosResults(c *gin.Context) {
	id := c.Param("id")

	//checa se o job ta finalizado
	tracker := config.GetRedisTracker()
	status, err := tracker.GetJobStatus(id)
	if err != nil {
		logg.Errorf("error checking job status: %v", err)
		response.SendError(c, http.StatusNotFound, fmt.Sprintf("error checking job status: %v", err))
		return
	}
	if status["status"] != "CONCLUIDO" {
		logg.Warning("job not completed yet")
		response.SendError(c, http.StatusAccepted, "job is being processed")
		return
	}
	//caso esteja, pegue o id e monte a saída
	result, err := getResults(id)
	if err != nil {
		logg.Errorf("could not get job results: %v", err)
		response.SendError(c, http.StatusNotFound, fmt.Sprintf("could not find job results: %v", err))
		return
	}

	response.SendSuccess(c, result)
}

func getResults(jobId string) (response.ResultResponse, error) {
	db := config.GetPostgresDB()
	var jobStatus job.JobStatus
	if err := db.First(&jobStatus, "job_id = ?", jobId).Error; err != nil {
		return response.ResultResponse{}, err
	}

	var cnpjs []job.JobCnpj
	if err := db.Where("job_id = ?", jobId).Find(&cnpjs).Error; err != nil {
		return response.ResultResponse{}, err
	}

	results := response.ResultResponse{
		JobId:        jobId,
		JobCreatedAt: jobStatus.JobCreatedAt,
		CompletedAt:  jobStatus.CompletedAt,
		Totals: response.TotalsStruct{
			Companies: jobStatus.Companies,
			Completed: jobStatus.Completed,
			Failed:    jobStatus.Failed,
		},
	}
	for _, jc := range cnpjs {
		var cmpny company.Company
		if err := db.First(&cmpny, "cnpj = ?", jc.Cnpj).Error; err != nil {
			return response.ResultResponse{}, err
		}

		var emails []response.EmailStruct
		db.Table("company_emails").Select("cnpj", "email", "valid").Where("cnpj = ?", jc.Cnpj).Find(&emails)

		var phones []response.PhoneStruct
		db.Table("company_phones").Select("cnpj", "phone", "valid").Where("cnpj = ?", jc.Cnpj).Find(&phones)

		results.Items = append(results.Items, response.ItemsStruct{
			Cnpj:              cmpny.Cnpj,
			RazaoSocial:       cmpny.Razao_social,
			Municipio:         cmpny.Municipio,
			Segmento:          types.Segmento(cmpny.Segmento),
			SituacaoCadastrao: types.Situation(cmpny.Situacao_cadastral),
			UpdatedAt:         cmpny.UpdatedAt,
			Telefones:         phones,
			Emails:            emails,
		})
	}

	return results, nil
}
