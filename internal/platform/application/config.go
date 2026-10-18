package application

import (
	"os"

	_os "github.com/kieranroneill/oarbit/internal/platform/os"
)

type Config struct {
	Environment string
	LogLevel    string
	Port        string
	Title       string
	Version     string
}

func NewConfig() Config {
	return Config{
		Environment: os.Getenv("ENVIRONMENT"),
		LogLevel:    _os.GetEnvWithDefault("LOG_LEVEL", "error"),
		Port:        _os.GetEnvWithDefault("PORT", "3000"),
		Title:       _os.GetEnvWithDefault("APP_TITLE", "Oarbit"),
		Version:     _os.GetEnvWithDefault("VERSION", "development"),
	}
}
