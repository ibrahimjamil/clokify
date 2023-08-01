package config

import (
	. "clokify/types"
	"os"
)

func EnvConfig() *Config {
	return &Config{
		Host:      os.Getenv("DB_HOST"),
		DBName:    os.Getenv("DB_NAME"),
		Port:      os.Getenv("DB_PORT"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		AppPort:   os.Getenv("APP_PORT"),
	}
}
