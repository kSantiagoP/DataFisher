package job

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type JobTracker struct {
	client *redis.Client
}

func NewTracker(redisURL string) (*JobTracker, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	return &JobTracker{
		client: redis.NewClient(opt),
	}, nil
}

// Publica novo job
func (t *JobTracker) CreateJob(jobId string, cnpj []string) error {
	ctx := context.Background()
	//Job structure:
	//Key: job:{id} -> {status, progress, total}
	err := t.client.HSet(ctx, "job:"+jobId,
		"status", "pending",
		"progress", 0,
		"total", len(cnpj),
	).Err()
	if err != nil {
		return fmt.Errorf("error creating job %s:%v", jobId, err)
	}

	for _, cnpj := range cnpj {
		t.client.SAdd(ctx, "job:"+jobId+":cnpjs", cnpj)
	}

	t.client.Expire(ctx, "job:"+jobId, 24*time.Hour)

	return nil
}

func (t *JobTracker) UpdateProgress(jobId string, current int) error {
	return t.client.HSet(context.Background(), "job:"+jobId,
		"progress", current,
		"status", "processing",
	).Err()
}

// GetJobStatus retorna o status, progresso e total de CNPJs
func (t *JobTracker) GetJobStatus(jobID string) (map[string]interface{}, error) {
	ctx := context.Background()

	// Busca os campos principais do hash
	result, err := t.client.HGetAll(ctx, "job:"+jobID).Result()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("job not found")
	}

	// Calcula porcentagem de conclusão
	progress, _ := strconv.Atoi(result["progress"])
	total, _ := strconv.Atoi(result["total"])
	percentage := 0
	if total > 0 {
		percentage = (progress * 100) / total
	}

	return map[string]interface{}{
		"status":     result["status"],
		"progress":   progress,
		"total":      total,
		"percentage": percentage,
	}, nil
}
