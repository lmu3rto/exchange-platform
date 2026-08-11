package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
	HttpAddr    string
}

func Load() (Config, error) {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		return Config{}, fmt.Errorf("db url is not set")
	}

	httpAddr := os.Getenv("HTTP_ADDR")

	if httpAddr == "" {
		httpAddr = ":8080"
	}

	return Config{
		DatabaseURL: dbURL,
		HttpAddr:    httpAddr,
	}, nil

}
