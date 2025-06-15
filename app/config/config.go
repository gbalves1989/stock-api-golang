package config

import (
	"os"
)

type Config struct {
	Port string
}

var AppConfig Config

func LoadConfig() {
	AppConfig.Port = getEnvDefault("PORT", "8080")
}

func getEnvDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}