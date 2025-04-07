package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing with system env...")
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); 
	exists {
		return value
	}
	return fallback
}
