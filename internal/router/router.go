package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/logger"
)

func Intialize() {
	logger := logger.NewLogger("router")
	router := gin.Default()
	router.GET("/ping", ping)

	logger.Debug("Server listening and running on port 8080")
	router.Run(":8080")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
