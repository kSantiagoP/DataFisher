package companyPhone

import "gorm.io/gorm"

type CompanyPhone struct {
	gorm.Model
	Cnpj  string
	Phone string
	Valid bool
}
