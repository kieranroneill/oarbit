package redis

import (
	"os"
	"strconv"
)

type Config struct {
	DB       int
	Host     string
	Password string
	Port     string
	Username string
}

func NewConfig() Config {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return Config{
		DB:       db,
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Port:     os.Getenv("REDIS_PORT"),
		Username: os.Getenv("REDIS_USERNAME"),
	}
}
