package main

import (
	"log"

	"first-go-project/configs"
	"first-go-project/internal/infrastructure/database"
	"first-go-project/internal/infrastructure/database/models"
)

func main() {
	config := configs.Load()

	db, err := database.NewPostgresDB(&config.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	users := []models.User{
		{
			Email:    "john.doe@example.com",
			Username: "johndoe",
			Password: "hashed_password_123",
			FullName: "John Doe",
			Phone:    "+1234567890",
			Address:  "123 Main Street",
			City:     "New York",
			Province: "NY",
			ZipCode:  "10001",
		},
		{
			Email:    "jane.smith@example.com",
			Username: "janesmith",
			Password: "hashed_password_456",
			FullName: "Jane Smith",
			Phone:    "+1987654321",
			Address:  "456 Oak Avenue",
			City:     "Los Angeles",
			Province: "CA",
			ZipCode:  "90001",
		},
		{
			Email:    "bob.wilson@example.com",
			Username: "bobwilson",
			Password: "hashed_password_789",
			FullName: "Bob Wilson",
			Phone:    "+1555123456",
			Address:  "789 Pine Road",
			City:     "Chicago",
			Province: "IL",
			ZipCode:  "60601",
		},
	}

	for _, user := range users {
		if err := db.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to create user %s: %v", user.Email, err)
		} else {
			log.Printf("✓ Created user: %s (ID: %d)", user.Email, user.ID)
		}
	}

	log.Println("✓ Seeding completed!")
}
