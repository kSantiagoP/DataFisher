package data_api

import "math/rand"

type PhoneValid struct {
	Valid bool   `json:"valid"`
	Phone string `json:"phone"`
}

func ValidatePhone(phone []string) []PhoneValid {
	//latency
	var results []PhoneValid
	for _, email := range phone {
		isValid := rand.Intn(2) == 1

		results = append(results, PhoneValid{
			Valid: isValid,
			Phone: email,
		})
	}
	return results
}
