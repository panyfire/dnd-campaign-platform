package config

type Config struct {
	AppName string
	AppEnv  string

	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string

	TelegramToken string
}
