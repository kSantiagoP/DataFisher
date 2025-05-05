package data_api

import (
	"time"
)

type PhoneValid struct {
	Valid bool   `json:"valid"`
	Phone string `json:"phone"`
}

func ValidatePhone(phones []string) []PhoneValid {
	time.Sleep(100 * time.Millisecond)
	var results []PhoneValid
	for _, phone := range phones {
		isValid := len(phone)%2 == 0

		results = append(results, PhoneValid{
			Valid: isValid,
			Phone: phone,
		})
	}
	return results
}
