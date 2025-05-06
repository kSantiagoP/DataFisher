package main

import (
	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/data_api"
	jobProcessing "github.com/kSantiagoP/DataFisher/internal/job_processing"
	"github.com/kSantiagoP/DataFisher/internal/logger"
)

func main() {
	logger := logger.NewLogger("worker")
	logger.Info("Worker started")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	err = config.InitDatabase()
	if err != nil {
		logger.Errorf("Error initializing database: %v", err)
		return
	}

	err = data_api.Init()
	if err != nil {
		logger.Errorf("Error connecting with dataApi: %v", err)
		return
	}

	err = waitForJobs()
	if err != nil {
		logger.Errorf("Error at job consumer: %v", err)
		return
	}
}

func waitForJobs() error {
	logger := logger.NewLogger("worker")
	queue := config.GetRabbitQueue()

	msgs, err := queue.Channel.Consume(
		"enrichment_jobs", // queue
		"",                // consumer
		false,             // auto-ack (importante: false para confirmação manual)
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		return err
	}

	logger.Info("Worker started. Waiting for jobs...")

	for msg := range msgs {
		job := msg.Body

		logger.Infof("Processing job: %v\n", string(job))

		//process
		err := jobProcessing.ProcessMessage(job)
		if err != nil {
			return err
		}

		msg.Ack(false)
	}

	return nil
}
