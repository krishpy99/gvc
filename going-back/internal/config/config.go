package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Config loaded")

	return &Config{Port: port}, nil
}
