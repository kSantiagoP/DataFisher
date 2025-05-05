package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/logger"
)

var logg *logger.Logger

func Intialize() {
	logg = logger.NewLogger("router")
	router := gin.Default()
	initializeRoutes(router)

	logg.Debug("Server listening and running on port 8080")
	router.Run(":8080")
}
