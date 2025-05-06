package config

import (
	"github.com/go-redis/redis/v8"
)

func initializeRedis() (*JobTracker, error) {
	redisURL := "redis://redis:6379"
	return newTracker(redisURL)
}

func newTracker(redisURL string) (*JobTracker, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	return &JobTracker{
		client: redis.NewClient(opt),
	}, nil
}
