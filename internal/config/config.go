package config

import (
	"fmt"

	"github.com/kSantiagoP/DataFisher/internal/logger"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	tracker *JobTracker
	queue   *Queue
	logg    *logger.Logger
)

func Init() error {
	var err error
	logg = logger.NewLogger("config")

	tracker, err = initializeRedis()
	if err != nil {
		return fmt.Errorf("error initializing tracker: %v", err)
	}

	queue, err = initializeRabbitMQ(6)
	if err != nil {
		return fmt.Errorf("error initializing queue: %v", err)
	}

	return nil
}

func InitDatabase() error {
	var err error
	db, err = initializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	return nil
}

func GetPostgresDB() *gorm.DB {
	return db
}

func GetRedisTracker() *JobTracker {
	return tracker
}

func GetRabbitQueue() *Queue {
	return queue
}
