package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_PASSWORD string
	DB_USER     string
	DATABASE    string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env... %v", err)
	}

	cfg := &Config{
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "3306"),
		DB_USER:     getEnv("DB_USER", "root"),
		DB_PASSWORD: getEnv("DB_PASSWORD", ""),
		DATABASE:    getEnv("DATABASE", ""),
	}

	// Required values check
	if cfg.DATABASE == "" {
		return nil, fmt.Errorf("missing required env: DATABASE")
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		log.Printf("[config] %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}

	return val
}
