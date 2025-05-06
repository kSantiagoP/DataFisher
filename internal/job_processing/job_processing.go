package jobProcessing

import (
	"encoding/json"
	"fmt"

	"github.com/kSantiagoP/DataFisher/internal/job_processing/handler"
	"github.com/kSantiagoP/DataFisher/internal/types"
)

type JobRequest struct {
	JobId     string   `json:"jobId"`
	Cnpjs     []string `json:"cnpjs"`
	Operation int      `json:"operation"`
}

func ProcessMessage(job []byte) error {
	request := JobRequest{}
	err := json.Unmarshal(job, &request)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	switch request.Operation {
	case int(types.ENRICH):
		return handler.EnrichCnpj(job)

	case int(types.CONSULT):
		return handler.ConsultJob(job)

	case int(types.RESULT):
		return nil
	}

	return nil
}
