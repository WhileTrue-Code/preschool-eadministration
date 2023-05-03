package config

import "os"

type Config struct {
	HealthcareServicePort string
	HealthcareDBPort      string
	HealthcareDBHost      string
}

func NewConfig() *Config {
	return &Config{
		HealthcareServicePort: os.Getenv("HEALTHCARE_SERVICE_PORT"),
		HealthcareDBHost:      os.Getenv("HEALTHCARE_DB_HOST"),
		HealthcareDBPort:      os.Getenv("HEALTHCARE_DB_PORT"),
	}
}
