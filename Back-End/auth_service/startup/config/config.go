package config

import "os"

type Config struct {
	AuthServicePort string
	AuthDBPort      string
	AuthDBHost      string
}

func NewConfig() *Config {
	return &Config{
		AuthServicePort: os.Getenv("AUTH_SERVICE_PORT"),
		AuthDBHost:      os.Getenv("AUTH_DB_HOST"),
		AuthDBPort:      os.Getenv("AUTH_DB_PORT"),
	}
}
