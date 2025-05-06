package companyEmail

import "gorm.io/gorm"

type CompanyEmail struct {
	gorm.Model
	Cnpj  string
	Email string
	Valid bool
}
