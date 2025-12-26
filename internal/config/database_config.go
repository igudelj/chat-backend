package config

import (
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func LoadDatabaseConfig() DatabaseConfig {
	config := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if config.Host == "" || config.Port == "" || config.Name == "" || config.User == "" {
		panic("database configuration is incomplete")
	}

	return config
}
