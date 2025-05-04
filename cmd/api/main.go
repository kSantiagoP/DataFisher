package main

import (
	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/data_api"
	"github.com/kSantiagoP/DataFisher/internal/logger"
	"github.com/kSantiagoP/DataFisher/internal/router"
)

func main() {
	logger := logger.NewLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	err = data_api.Init()
	if err != nil {
		logger.Errorf("Error connecting with dataApi: %v", err)
		return
	}

	router.Intialize()
}
