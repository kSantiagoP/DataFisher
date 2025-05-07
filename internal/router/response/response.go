package response

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kSantiagoP/DataFisher/internal/types"
)

type TotalsStruct struct {
	Companies int `json:"companies"`
	Completed int `json:"completed"`
	Failed    int `json:"failed"`
	Pending   int `json:"pending,omitempty"`
}

type ConsultaResponse struct {
	JobId      string       `json:"job_id"`
	Status     string       `json:"status"`
	Progress   float64      `json:"progress"`
	Totals     TotalsStruct `json:"totals"`
	LastUpdate *time.Time   `json:"last_update,omitempty"`
	ResultsUrl string       `json:"results_url,omitempty"`
}

type PhoneStruct struct {
	Numero string `gorm:"column:phone" json:"numero"`
	Valido bool   `gorm:"column:valido" json:"valido"`
}
type EmailStruct struct {
	Email  string `gorm:"column:email" json:"email"`
	Valido bool   `gorm:"column:valido" json:"valido"`
}

type ItemsStruct struct {
	Cnpj              string          `json:"cnpj"`
	RazaoSocial       string          `json:"razao_social"`
	Municipio         string          `json:"municipio"`
	Segmento          types.Segmento  `json:"segmento"`
	SituacaoCadastrao types.Situation `json:"situacao_cadastral"`
	UpdatedAt         time.Time       `json:"updated_at"`
	Telefones         []PhoneStruct   `json:"telefones"`
	Emails            []EmailStruct   `json:"emails"`
}

type ResultResponse struct {
	JobId        string        `json:"job_id"`
	JobCreatedAt string        `json:"created_at"`
	CompletedAt  string        `json:"completed_at"`
	Totals       TotalsStruct  `json:"totals"`
	Items        []ItemsStruct `json:"items"`
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
