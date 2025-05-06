package config

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

// Publica novo job
func (t *JobTracker) CreateJob(jobId string, cnpj []string) error {
	ctx := context.Background()
	//Job structure:
	//Key: job:{id} -> {status, progress, total}
	err := t.client.HSet(ctx, "job:"+jobId,
		"status", "EM_ANDAMENTO",
		"progress", 0,
		"total", len(cnpj),
		"failed", 0,
		"pending", len(cnpj),
		"created_at", time.Now().Format(time.RFC3339),
		"last_update", time.Now().Format(time.RFC3339),
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

func (t *JobTracker) IncrementProgress(jobId string) error {
	ctx := context.Background()

	_, err := t.client.HIncrBy(ctx, "job:"+jobId, "progress", 1).Result()
	if err != nil {
		return err
	}

	_, err = t.client.HIncrBy(ctx, "job:"+jobId, "pending", -1).Result()
	if err != nil {
		return err
	}

	_, err = t.client.HSet(ctx, "job:"+jobId,
		"status", "EM_ANDAMENTO",
		"last_update", time.Now().Format(time.RFC3339),
	).Result()
	return err
}

func (t *JobTracker) CompleteJob(jobId string) error {
	ctx := context.Background()
	_, err := t.client.HDel(ctx, "job:"+jobId, "last_update").Result()
	if err != nil {
		return err
	}
	return t.client.HSet(ctx, "job:"+jobId,
		"status", "CONCLUIDO",
		"completed_at", time.Now().Format(time.RFC3339),
	).Err()
}

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

	failed, err := t.GetFailedCount(jobID)
	if err != nil {
		return nil, err
	}

	// Calcula porcentagem de conclusão
	progress, _ := strconv.Atoi(result["progress"])
	total, _ := strconv.Atoi(result["total"])
	pending, _ := strconv.Atoi(result["pending"])
	percentage := 0.0
	if total > 0 {
		percentage = float64(progress) / float64(total)
	}

	return map[string]interface{}{
		"job_id":    jobID,
		"status":    result["status"],
		"progress":  percentage,
		"companies": total,
		"completed": total - int(failed),
		"failed":    int(failed),
		"pending":   pending - int(failed),
	}, nil
}

func (t *JobTracker) MarkFailedCNPJ(jobId, cnpj string, reason error) error {
	ctx := context.Background()

	err := t.client.SAdd(ctx, "job:"+jobId+":failed", cnpj).Err()
	if err != nil {
		return err
	}

	return t.client.HSet(ctx, "job:"+jobId+":fail_reasons",
		cnpj, reason.Error(),
		"last_update", time.Now().Format(time.RFC3339),
	).Err()
}

func (t *JobTracker) StartJob(jobId string) error {
	return t.client.HSet(context.Background(), "job:"+jobId,
		"status", "processing",
		"started_at", time.Now().Format(time.RFC3339),
	).Err()
}

func (t *JobTracker) GetFailedCount(jobId string) (int64, error) {
	return t.client.SCard(context.Background(), "job:"+jobId+":failed").Result()
}

func (t *JobTracker) PartiallyCompleteJob(jobId string, failedCount int64) error {
	ctx := context.Background()
	_, err := t.client.HDel(ctx, "job:"+jobId, "last_update").Result()
	if err != nil {
		return err
	}
	return t.client.HSet(ctx, "job:"+jobId,
		"status", "CONCLUIDO",
		"completed_at", time.Now().Format(time.RFC3339),
		"failed", failedCount,
	).Err()
}
