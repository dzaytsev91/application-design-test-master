package config

import (
	"os"
)

// Config service config
type Config struct {
	Host string
	Port string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "8080"),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
