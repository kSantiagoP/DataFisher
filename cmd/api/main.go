package main

import (
	"github.com/kSantiagoP/DataFisher/internal/config"
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

	router.Intialize()
}
