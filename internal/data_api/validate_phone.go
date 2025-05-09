package data_api

import (
	"fmt"
	"strconv"
	"time"

	companyPhone "github.com/kSantiagoP/DataFisher/internal/model/company_phone"
)

type PhoneValid struct {
	Valid bool   `json:"valid"`
	Phone string `json:"phone"`
}

func ValidatePhone(phones []string, cnpj string) []companyPhone.CompanyPhone {
	time.Sleep(100 * time.Millisecond)
	var results []companyPhone.CompanyPhone
	for _, phone := range phones {
		_, isValid, _ := GetLastCharAndCheckEven(phone)

		results = append(results, companyPhone.CompanyPhone{
			Cnpj:  cnpj,
			Valid: isValid,
			Phone: phone,
		})
	}
	return results
}

func GetLastCharAndCheckEven(s string) (int, bool, error) {
	if len(s) == 0 {
		return 0, false, fmt.Errorf("string vazia")
	}

	lastChar := s[len(s)-1:]
	num, err := strconv.Atoi(lastChar)
	if err != nil {
		return 0, false, fmt.Errorf("último caractere não é número")
	}

	return num, num%2 == 0, nil
}
