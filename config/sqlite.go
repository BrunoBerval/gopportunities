package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/BrunoBerval/gopportunities/schemas"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger()
	dbPath := "./db/main.db"

	// Checando se o banco já existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database not found, creating...")
		// Cria o banco de dados
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	// Criação da banco de dados
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migração
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Retorna o banco
	return db, nil
}
