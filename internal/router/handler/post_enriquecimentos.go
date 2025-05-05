package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/router/request"
	"github.com/kSantiagoP/DataFisher/internal/router/response"
)

func PostEnriquecimentos(c *gin.Context) {
	request := request.PostJobRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logg.Errorf("validation error: %v", err.Error())
		response.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	jobId := uuid.New().String()

	//entrega json com cnpj (o request) para o rabbitmq e trabalha daí
	tracker := config.GetRedisTracker()
	if err := tracker.CreateJob(jobId, request.Cnpjs); err != nil {
		logg.Errorf("job creation error: %v", err)
		response.SendError(c, http.StatusInternalServerError, "failed to create job")
		return
	}

	queue := config.GetRabbitQueue()
	if err := queue.Publish(jobId, request.Cnpjs); err != nil {
		logg.Errorf("publish error: %v", err)
		response.SendError(c, http.StatusInternalServerError, "failed to publish job")
		return
	}

	response.SendSuccess(c, gin.H{
		"job_id":          jobId,
		"total_companies": len(request.Cnpjs),
		"created_at":      time.Now(),
	})
}
