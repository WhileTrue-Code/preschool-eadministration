package config

import "os"

type Config struct {
	SERVICE_PORT string
	DB_HOST      string
	DB_PORT      string
}

func NewConfig() *Config {
	return &Config{
		SERVICE_PORT: os.Getenv("SERVICE_PORT"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_PORT:      os.Getenv("DB_PORT"),
	}
}
