// Package config
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("config: ", err)
	}
	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("config : DATABASE_URL not set")
	}

	return cfg
}
