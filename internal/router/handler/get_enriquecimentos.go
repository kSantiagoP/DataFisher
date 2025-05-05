package handler

import "github.com/gin-gonic/gin"

func GetEnriquecimentos(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "alo frequesia",
		"param":   id,
	})
}
