package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/logger"
)

func Intialize() {
	logger := logger.NewLogger("router")
	router := gin.Default()
	initializeRoutes(router)

	logger.Debug("Server listening and running on port 8080")
	router.Run(":8080")
}
