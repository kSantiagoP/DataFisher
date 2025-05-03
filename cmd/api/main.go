package main

import (
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
}
