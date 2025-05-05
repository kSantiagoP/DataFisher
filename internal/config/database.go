package config

import (
	"github.com/kSantiagoP/DataFisher/internal/logger"
	"github.com/kSantiagoP/DataFisher/internal/model/company"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := logger.NewLogger("postgres")
	dbUrl := "postgres://postgres:postgres@postgres:5432/datafisher_db?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logger.Errorf("error initializing database: %v", err)
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
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}
	logger.Debug("Postgres online.")

	return db, nil
}
