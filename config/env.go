package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV::failed loading .env file -> err: %v", err)
	}
}
