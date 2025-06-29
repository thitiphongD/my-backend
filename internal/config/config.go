package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Port             string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPass           string
	DBName           string
	DBSSLMode        string
	DBChannelBinding string
	JWTSecret        string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists (for development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		Port:             getEnv("PORT", "8080"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPass:           getEnv("DB_PASS", "password"),
		DBName:           getEnv("DB_NAME", "mydb"),
		DBSSLMode:        getEnv("DB_SSL_MODE", "disable"),
		DBChannelBinding: getEnv("DB_CHANNEL_BINDING", ""),
		JWTSecret:        getEnv("JWT_SECRET", "your-secret-key"),
	}

	// Validate required configuration
	if config.JWTSecret == "your-secret-key" {
		log.Println("WARNING: Using default JWT secret. Please set JWT_SECRET environment variable in production")
	}

	return config
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
