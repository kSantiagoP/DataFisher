package handler

import (
	"github.com/kSantiagoP/DataFisher/internal/logger"
)

var (
	logg *logger.Logger
)

func InitializeHandler() {
	logg = logger.NewLogger("handler")
}
