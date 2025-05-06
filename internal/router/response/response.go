package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type TotalsStruct struct {
	Companies int `json:"companies"`
	Completed int `json:"completed"`
	Failed    int `json:"failed"`
	Pending   int `json:"pending"`
}

type ConsultaResponse struct {
	JobId      string       `json:"job_id"`
	Status     string       `json:"status"`
	Progress   float64      `json:"progress"`
	Totals     TotalsStruct `json:"totals"`
	LastUpdate *time.Time   `json:"last_update,omitempty"`
}

func SendSuccess(c *gin.Context, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(200, data)
}

func SendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
