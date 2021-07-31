package config

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file.")
	}
}
