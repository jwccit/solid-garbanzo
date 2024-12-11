package config

import (
	"os"
)

var EnvConfig map[string]string

func LoadConfig() {
	EnvConfig = make(map[string]string)
	EnvConfig["DB_HOST"] = getEnv("DB_HOST", "localhost")
	EnvConfig["DB_USER"] = getEnv("DB_USER", "root")
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
