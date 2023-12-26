package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	DbUser = os.Getenv("MYSQL_USER")
	DbPassword = os.Getenv("MYSQL_PASSWORD")
	DbHost = os.Getenv("MYSQL_HOST")
	DbPort = os.Getenv("MYSQL_PORT")
	DbName = os.Getenv("MYSQL_DATABASE")
}
