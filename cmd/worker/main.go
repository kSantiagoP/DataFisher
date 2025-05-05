package main

import "github.com/kSantiagoP/DataFisher/internal/logger"

func main() {
	logger := logger.NewLogger("worker")
	logger.Debug("SALVE SALVE YODINHA FELIZ ANIVERSÁRIO")
}
