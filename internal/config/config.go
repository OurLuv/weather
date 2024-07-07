package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User         string
	Password     string
	DatabaseName string
	KEY          string
}

func MustLoad() *Config {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	var exists bool
	cfg.User, exists = os.LookupEnv("DB_USER")
	if !exists {
		panic("user is empty")
	}
	cfg.Password, exists = os.LookupEnv("DB_PASSWORD")
	if !exists {
		panic("password  is empty")
	}
	cfg.DatabaseName, exists = os.LookupEnv("DB_NAME")
	if !exists {
		panic("database name is empty")
	}
	cfg.KEY, exists = os.LookupEnv("API_KEY")
	if !exists {
		panic("KEY is empty")
	}

	return &cfg
}
