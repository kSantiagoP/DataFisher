package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/model/job"
	"github.com/kSantiagoP/DataFisher/internal/router/response"
	"gorm.io/gorm"
)

func GetEnriquecimentos(c *gin.Context) {
	id := c.Param("id")
	var err error
	tracker := config.GetRedisTracker()
	result, _ := tracker.GetJobStatus(id)
	if len(result) == 0 {
		result, err = getJobStatusDb(id)

		if err != nil {
			response.SendError(c, http.StatusInternalServerError, "could not retrieve status")
			return
		}
		if len(result) == 0 {
			response.SendError(c, http.StatusNotFound, "no job registered with provided id")
			return
		}
	}

	totals := response.TotalsStruct{
		Companies: result["companies"].(int),
		Completed: result["completed"].(int),
		Failed:    result["failed"].(int),
		Pending:   result["pending"].(int),
	}
	res := response.ConsultaResponse{
		JobId:    result["job_id"].(string),
		Status:   result["status"].(string),
		Progress: result["progress"].(float64),
		Totals:   totals,
	}

	if lastResponse, ok := result["last_update"].(time.Time); ok {
		res.LastUpdate = &lastResponse
	}
	if result["status"].(string) == "CONCLUIDO" {
		res.ResultsUrl = "http://localhost:8080/enriquecimentos/" + id + "/results"
	}
	response.SendSuccess(c, res)
}

func getJobStatusDb(jobId string) (map[string]interface{}, error) {
	db := config.GetPostgresDB()
	jobStatus := job.JobStatus{}
	if err := db.Where("job_id = ?", jobId).First(&jobStatus).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, nil
	}

	return map[string]interface{}{
		"job_id":    jobStatus.JobId,
		"status":    jobStatus.Status,
		"progress":  jobStatus.Progress,
		"companies": jobStatus.Companies,
		"completed": jobStatus.Completed,
		"failed":    jobStatus.Failed,
		"peding":    jobStatus.Pending,
	}, nil
}
