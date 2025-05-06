package handler

type JobRequest struct {
	JobId     string   `json:"jobId"`
	Cnpjs     []string `json:"cnpjs"`
	Operation int      `json:"operation"`
}
