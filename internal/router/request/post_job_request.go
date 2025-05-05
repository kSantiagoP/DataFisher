package request

import (
	"fmt"
	"regexp"
	"strconv"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type %s) is required", name, typ)
}

type PostJobRequest struct {
	Cnpjs []string `json:"cnpjs"`
}

func (r *PostJobRequest) Validate() error {
	if len(r.Cnpjs) == 0 {
		return errParamIsRequired("cnpj", "string array")
	}

	for _, cnpj := range r.Cnpjs {
		if err := validateCNPJ(cnpj); err != nil {
			return err
		}
	}

	return nil
}

func validateCNPJ(cnpj string) error {
	if !isOnlyDigits(cnpj) {
		return fmt.Errorf("CNPJ %s must have only 14 digits", cnpj)
	}

	if len(cnpj) != 14 {
		return fmt.Errorf("CNPJ %s must contain exactly 14 digits", cnpj)
	}

	if !validCheckDigits(cnpj) {
		return fmt.Errorf("CNPJ %s is invalid", cnpj)
	}

	return nil
}

func isOnlyDigits(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}

func validCheckDigits(cnpj string) bool {
	// Peso para cálculo do primeiro dígito verificador
	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	// Peso para cálculo do segundo dígito verificador
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calcula primeiro dígito verificador
	sum := 0
	for i := range 12 {
		digit, _ := strconv.Atoi(string(cnpj[i]))
		sum += digit * weights1[i]
	}
	rest := sum % 11
	digit1 := 0
	if rest >= 2 {
		digit1 = 11 - rest
	}

	// Calcula segundo dígito verificador
	sum = 0
	for i := range 13 {
		digit, _ := strconv.Atoi(string(cnpj[i]))
		sum += digit * weights2[i]
	}
	rest = sum % 11
	digit2 := 0
	if rest >= 2 {
		digit2 = 11 - rest
	}

	// Verifica se os dígitos calculados batem com os informados
	actualDigit1, _ := strconv.Atoi(string(cnpj[12]))
	actualDigit2, _ := strconv.Atoi(string(cnpj[13]))

	return digit1 == actualDigit1 && digit2 == actualDigit2
}
