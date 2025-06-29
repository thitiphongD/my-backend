package database

import (
	"fmt"
	"log"

	"github.com/thitiphongD/my-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDatabase initializes database connection using config
func ConnectDatabase() {
	cfg := config.LoadConfig()

	// สร้าง connection string จาก config parameters
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	// เพิ่ม channel_binding หากมีการกำหนดค่า
	if cfg.DBChannelBinding != "" {
		connectionString += fmt.Sprintf(" channel_binding=%s", cfg.DBChannelBinding)
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
