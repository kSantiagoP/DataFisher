package config

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func initializeRabbitMQ(maxAttempts int) (*Queue, error) {
	rabbitmqURL := "amqp://guest:guest@rabbitmq:5672/"

	for i := range maxAttempts {
		q, err := newQueue(rabbitmqURL)
		if err == nil {
			return q, nil
		}

		fmt.Printf("Attempt %d/%d: RabbitMQ not ready (%v)\n", i+1, maxAttempts, err)
		time.Sleep(exponentialBackoff(i))
	}

	return nil, fmt.Errorf("RabbitMQ could not be initialized after %d attempts", maxAttempts)
}

func exponentialBackoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * time.Second
}

type Queue struct {
	Connec  *amqp.Connection
	Channel *amqp.Channel
}

// Cria nova fila de jobs.
func newQueue(rabbitmqURL string) (*Queue, error) {
	connec, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	ch, err := connec.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"jobs",   //name
		"direct", //type
		true,     //durable
		false,    //auto-deleted
		false,    //internal
		false,    //no-wait
		nil,      //arguments
	)
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"enrichment_jobs", // name
		true,              // durable (sobrevive a reinícios)
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Queue{Connec: connec, Channel: ch}, nil
}

type JobMessage struct {
	JobId string   `json:"jobId"`
	Cnpjs []string `json:"cnpjs"`
}

// Publica um job na fila
func (q *Queue) Publish(jobID string, cnpjs []string) error {
	message := JobMessage{
		JobId: jobID,
		Cnpjs: cnpjs,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = q.Channel.Publish(
		"jobs",       // exchange
		"enrichment", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // Persistente em disco
			ContentType:  "application/json",
			Body:         body,
			Timestamp:    time.Now(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	logg.Infof("Job %s published to RabbitMQ with %d CNPJs", jobID, len(cnpjs))

	return nil
}

func (r *Queue) Close() {
	r.Channel.Close()
	r.Connec.Close()
}
