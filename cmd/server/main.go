package main

import (
	"fmt"
	"log"
	"net/http"

	"first-go-project/configs"
	"first-go-project/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.Load()

	// Initialize database connection
	db, err := database.NewPostgresDB(&config.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Run migrations (add your models here when ready)
	// if err := db.AutoMigrate(&models.User{}); err != nil {
	//     log.Fatalf("Failed to run migrations: %v", err)
	// }

	router := gin.Default()

	// Health check endpoint
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Pong",
		})
	})

	// Start server
	addr := fmt.Sprintf(":%s", config.App.Port)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
