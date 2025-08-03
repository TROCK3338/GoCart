package database

import (
	"log"

	"GoCart/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gocart.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	log.Println("Database connection successful!")

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})
	if err != nil {
		log.Fatal("Failed to auto-migrate database schema!")
	}

	log.Println("Database migration successful!")
}