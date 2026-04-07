package config

import (
	"os"
)

type Config struct {
	Port  string
	DBURL string
}

func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/yeyen?sslmode=disable"
	}

	return Config{
		Port:  port,
		DBURL: dbURL,
	}
}
