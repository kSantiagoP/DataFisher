package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/router/response"
)

func GetEnriquecimentosResults(c *gin.Context) {
	id := c.Param("id")

	response.SendSuccess(c, gin.H{
		"message": "chegouResults",
		"id":      id,
	})
}
