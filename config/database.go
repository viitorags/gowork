package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/viitorags/gowork/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	logger = NewLogger("postgres")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Postgres opening error: ", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Works{})
	if err != nil {
		logger.Error("Postgres automigrate error: ", err)
		return nil, err
	}

	return db, nil

}
