package config

import "os"

type Config struct {
	RegistrarServicePort string
	RegistrarDBPort      string
	RegistrarDBHost      string
}

func NewConfig() *Config {
	return &Config{
		RegistrarServicePort: os.Getenv("REGISTRAR_SERVICE_PORT"),
		RegistrarDBHost:      os.Getenv("REGISTRAR_DB_HOST"),
		RegistrarDBPort:      os.Getenv("REGISTRAR_DB_PORT"),
	}
}
