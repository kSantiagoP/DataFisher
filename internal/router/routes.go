package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/router/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("")
	{
		v1.GET("/enriquecimentos/:id", handler.GetEnriquecimentos)
		v1.GET("/enriquecimentos/:id/results", handler.GetEnriquecimentosResults)
		v1.POST("/enriquecimentos", handler.PostEnriquecimentos)
	}
}
