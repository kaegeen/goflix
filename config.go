package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	DBConnection  string
	JWTSecret     string
}

func LoadConfig() (*Config, error) {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = ":8080"
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	if dbConnection == "" {
		dbConnection = "goflix.db"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "supersecretkey"
	}

	return &Config{
		ServerAddress: serverAddress,
		DBConnection:  dbConnection,
		JWTSecret:     jwtSecret,
	}, nil
}
