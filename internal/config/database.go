package config

import (
	"github.com/kSantiagoP/DataFisher/internal/model/company"
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

	//migrate schema
	if !db.Migrator().HasTable(&company.Company{}) {
		err = db.AutoMigrate(&company.Company{})
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		logg.Errorf("postgres automigration error: %v", err)
		return nil, err
	}
	logg.Debug("Postgres online.")

	return db, nil
}
