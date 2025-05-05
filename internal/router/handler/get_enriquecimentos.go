package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/router/response"
)

func GetEnriquecimentos(c *gin.Context) {
	id := c.Param("id")

	response.SendSuccess(c, gin.H{
		"message": "alo frequesia",
		"param":   id,
	})
}
