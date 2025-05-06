package job

import (
	"gorm.io/gorm"
)

type JobCnpj struct {
	gorm.Model
	JobId string
	Cnpj  string
}

type JobStatus struct {
	gorm.Model
	JobId        string
	Status       string
	Progress     float64
	Companies    int
	Completed    int
	Failed       int
	Pending      int
	JobCreatedAt string
	CompletedAt  string
}
