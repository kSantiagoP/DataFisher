package handler

import (
	"github.com/kSantiagoP/DataFisher/internal/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeHandler() {
	db = config.GetPostgresDB()
}
