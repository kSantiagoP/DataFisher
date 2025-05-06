package config

import (
	"github.com/kSantiagoP/DataFisher/internal/model/company"
	companyEmail "github.com/kSantiagoP/DataFisher/internal/model/company_email"
	companyPhone "github.com/kSantiagoP/DataFisher/internal/model/company_phone"
	"github.com/kSantiagoP/DataFisher/internal/model/job"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() (*gorm.DB, error) {
	dbUrl := "postgres://postgres:postgres@postgres:5432/datafisher_db?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logg.Errorf("error initializing database: %v", err)
		return nil, err
	}

	logg.Debug("Postgres online.")

	return db, nil
}

func MigrateSchemas() error {
	dbUrl := "postgres://postgres:postgres@postgres:5432/datafisher_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logg.Errorf("error initializing database: %v", err)
		return err
	}

	if !db.Migrator().HasTable(&company.Company{}) {
		err = db.AutoMigrate(&company.Company{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&companyEmail.CompanyEmail{}) {
		err = db.AutoMigrate(&companyEmail.CompanyEmail{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&companyPhone.CompanyPhone{}) {
		err = db.AutoMigrate(&companyPhone.CompanyPhone{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&job.JobCnpj{}) {
		err = db.AutoMigrate(&job.JobCnpj{})
		if err != nil {
			return err
		}
	}

	return nil
}
