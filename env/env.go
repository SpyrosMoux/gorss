package env

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariable(variable string) string {
	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found, attempting to read variables from host environment")
	}

	variableValue := getEnvOrExit(variable)

	return variableValue
}

func getEnvOrExit(key string) string {
	value := os.Getenv(key)
	if value == "" {
		slog.Error("environment variable not set", "variable", key)
		os.Exit(1)
	}
	return value
}
