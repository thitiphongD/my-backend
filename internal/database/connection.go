package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDatabase initializes database connection using connection string
func ConnectDatabase() {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("DB_CONNECTION_STRING is not set in the .env file")
	}

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected successfully!")
	DB = database
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// AutoMigrate runs auto migration for given models
func AutoMigrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
