package handler

import (
	"github.com/gin-gonic/gin"
)

func GetEnriquecimentosResults(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "chegouResults",
		"id":      id,
	})
}
