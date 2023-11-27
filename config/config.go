package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error when loading .env file")
	}
	return os.Getenv(key)
}
