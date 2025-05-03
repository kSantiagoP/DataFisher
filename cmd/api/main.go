package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/logger"
	"github.com/kSantiagoP/DataFisher/internal/model/company"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := logger.NewLogger("main")

	dbUrl := "postgres://postgres:postgres@postgres:5432/datafisher_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logger.Errorf("error initializing database: %v", err)
		return
	}

	//migrate schema
	err = db.AutoMigrate(&company.Company{})

	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return
	}
	logger.Debug("Postgres online.")

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
