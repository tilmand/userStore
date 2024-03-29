package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("SQL_HOST")
	port := os.Getenv("SQL_PORT")
	database := os.Getenv("SQL_DATABASE")
	username := os.Getenv("SQL_USERNAME")
	password := os.Getenv("SQL_PWD")

	config := &Config{
		Host:     host,
		Port:     port,
		Database: database,
		Username: username,
		Password: password,
	}

	return config, nil
}
