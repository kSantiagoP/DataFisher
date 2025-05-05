package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/router/response"
)

func PostEnriquecimentos(c *gin.Context) {

	response.SendSuccess(c, gin.H{
		"message": "chegouPost",
	})
}
