package data_api

import (
	"time"

	companyEmail "github.com/kSantiagoP/DataFisher/internal/model/company_email"
)

type EmailValid struct {
	Valid bool   `json:"valid"`
	Email string `json:"email"`
}

func ValidateEmail(emails []string, cnpj string) []companyEmail.CompanyEmail {
	time.Sleep(100 * time.Millisecond)

	var results []companyEmail.CompanyEmail
	for _, email := range emails {
		isValid := len(email)%2 == 0

		results = append(results, companyEmail.CompanyEmail{
			Cnpj:  cnpj,
			Valid: isValid,
			Email: email,
		})
	}
	return results
}
