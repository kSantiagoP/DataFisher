package job

import "gorm.io/gorm"

type JobCnpj struct {
	gorm.Model
	JobId string
	Cnpj  string
}
