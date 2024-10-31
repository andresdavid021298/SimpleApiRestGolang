package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadProperty(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading file .env")
	}
	property := os.Getenv(key)
	return property
}
