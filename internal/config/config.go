package config

import (
	"log"
	"os"
)

// Config holds all configuration values
type Config struct {
	DBURL      string
	Port       string
	GeminiKey  string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	config := &Config{
		DBURL:     getEnvOrDefault("DB_URL", ""),
		Port:      getEnvOrDefault("PORT", "8080"),
		GeminiKey: getEnvOrDefault("GEMINI_API_KEY", ""),
	}

	// Validate required fields
	if config.DBURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}
	if config.GeminiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}

	return config
}

// getEnvOrDefault returns the value of the environment variable or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 