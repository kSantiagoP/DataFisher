package config

import (
	"fmt"

	"github.com/kSantiagoP/DataFisher/internal/job"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	tracker *job.JobTracker
)

func Init() error {
	var err error
	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}

	redisURL := "redis://redis:6379"
	tracker, err = job.NewTracker(redisURL)
	if err != nil {
		return fmt.Errorf("error initializing tracker: %v", err)
	}

	return nil
}

func GetPostgresDB() *gorm.DB {
	return db
}

func GetRedisTracker() *job.JobTracker {
	return tracker
}
