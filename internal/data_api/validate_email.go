package data_api

import (
	"math/rand"
)

type EmailValid struct {
	Valid bool   `json:"valid"`
	Email string `json:"email"`
}

func ValidateEmail(emails []string) []EmailValid {
	//latency

	var results []EmailValid
	for _, email := range emails {
		isValid := rand.Intn(2) == 1

		results = append(results, EmailValid{
			Valid: isValid,
			Email: email,
		})
	}
	return results
}
