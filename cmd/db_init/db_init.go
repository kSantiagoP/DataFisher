package main

import (
	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/logger"
)

func main() {
	logg := logger.NewLogger("db_init")
	err := config.MigrateSchemas()
	if err != nil {
		logg.Errorf("Error initializing database: %v", err)
		return
	}
}
