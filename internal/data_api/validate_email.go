package data_api

import (
	"time"
)

type EmailValid struct {
	Valid bool   `json:"valid"`
	Email string `json:"email"`
}

func ValidateEmail(emails []string) []EmailValid {
	time.Sleep(100 * time.Millisecond)

	var results []EmailValid
	for _, email := range emails {
		isValid := len(email)%2 == 0

		results = append(results, EmailValid{
			Valid: isValid,
			Email: email,
		})
	}
	return results
}
