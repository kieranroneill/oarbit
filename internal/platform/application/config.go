package application

import (
	"os"
	"strings"

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

func (c Config) ParseMajorVersion() string {
	version := strings.Clone(c.Version)

	if i := strings.IndexAny(version, "-+"); i != -1 {
		version = version[:i]
	}

	major, _, _ := strings.Cut(version, ".")

	return major
}
