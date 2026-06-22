package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	err := godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		AppName:       os.Getenv("APP_NAME"),
		AppEnv:        os.Getenv("APP_ENV"),
		DBHost:        os.Getenv("DB_HOST"),
		DBName:        os.Getenv("DB_NAME"),
		DBUser:        os.Getenv("DB_USER"),
		DBPort:        port,
		DBPassword:    os.Getenv("DB_PASSWORD"),
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}

	return cfg, nil
}
