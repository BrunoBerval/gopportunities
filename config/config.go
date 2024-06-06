package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger() *Logger {
	// Inicia o Logger
	logger = NewLogger()
	return logger
}
