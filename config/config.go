package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSqlite()
	if err != nil {
		return fmt.Errorf("error initializing sql: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	// Inicia o Logger
	logger = NewLogger()
	return logger
}