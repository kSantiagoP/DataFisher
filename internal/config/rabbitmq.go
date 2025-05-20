package config

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func initializeRabbitMQ(maxAttempts int) (*Queue, error) {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

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
	JobId     string   `json:"jobId"`
	Cnpjs     []string `json:"cnpjs"`
	Operation int      `json:"operation"`
}

// Publica um job na fila
func (q *Queue) Publish(jobID string, cnpjs []string, operation int) error {
	message := JobMessage{
		JobId:     jobID,
		Cnpjs:     cnpjs,
		Operation: operation,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = q.Channel.ExchangeDeclare(
		"jobs",   // name
		"direct", // type
		true,     // durable
		false,    // auto-delete
		false,    // internal
		false,    // no-wait
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to exchange: %v", err)
	}
	// vinculate queue to exchange
	err = q.Channel.QueueBind(
		"enrichment_jobs", // queue name
		"enrichment",      // routing key
		"jobs",            // exchange name
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind: %v", err)
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
