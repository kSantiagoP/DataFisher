package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := logger.NewLogger("main")

	dbUrl := "postgres://postgres:postgres@postgres:5432/datafisher_db?sslmode=disable"
	_, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logger.Errorf("error initializing database: %v", err)
		return
	}
	logger.Debug("Postgres online.")

	router := gin.Default()
	router.GET("/ping", ping)
	router.Run(":8080") // listen and serve on localhost:4000
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
