package response

import (
	"time"

	"github.com/gin-gonic/gin"
	companyEmail "github.com/kSantiagoP/DataFisher/internal/model/company_email"
	companyPhone "github.com/kSantiagoP/DataFisher/internal/model/company_phone"
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
}

type PhoneStruct struct {
	Numero string `json:"numero"`
	Valido bool   `json:"valido"`
}
type EmailStruct struct {
	Email  string `json:"email"`
	Valido bool   `json:"valido"`
}

type ItemsStruct struct {
	Cnpj              string                      `json:"cnpj"`
	RazaoSocial       string                      `json:"razao_social"`
	Municipio         string                      `json:"municipio"`
	Segmento          types.Segmento              `json:"segmento"`
	SituacaoCadastrao types.Situation             `json:"situacao_cadastral"`
	UpdatedAt         time.Time                   `json:"updated_at"`
	Telefones         []companyPhone.CompanyPhone `json:"telefones"`
	Emails            []companyEmail.CompanyEmail `json:"emails"`
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
