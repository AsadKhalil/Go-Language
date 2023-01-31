package config

import (
	"os"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     string
}

func NewConfig() *Config {
	return &Config{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
	}
}
