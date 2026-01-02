package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
	err    error
)

func Init() error {

	db, err = InitializeDb()
	if err != nil {
		return fmt.Errorf("Error initialize database: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)

	return logger
}
