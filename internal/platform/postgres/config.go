package postgres

import (
	"os"

	_os "github.com/kieranroneill/oarbit/internal/platform/os"
)

type Config struct {
	Host     string
	Name     string
	Password string
	Port     string
	SSLMode  string
	User     string
}

func NewConfig() Config {
	return Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Name:     os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SSLMode:  _os.GetEnvWithDefault("POSTGRES_SSL_MODE", "enable"),
		User:     os.Getenv("POSTGRES_USER"),
	}
}
