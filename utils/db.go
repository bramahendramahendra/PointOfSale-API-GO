package utils

import (
	"apps/config"
	"apps/models"
	"log"

	"github.com/joho/godotenv"
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	config.DB.AutoMigrate(&models.User{})
}
