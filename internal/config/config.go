package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            int
	Environment     string
	DatabaseURL     string
	OpenAIKey       string
	ShutdownTimeout time.Duration
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	timeout, err := time.ParseDuration(os.Getenv("SHUTDOWN_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("invalid shutdown timeout: %w", err)
	}

	return &Config{
		Port:            getEnvInt("PORT", 8080),
		Environment:     os.Getenv("ENVIRONMENT"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		OpenAIKey:       os.Getenv("OPENAI_API_KEY"),
		ShutdownTimeout: timeout,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
	}
	return defaultValue
}
