package main

import (
	"fmt"
	"log"
	"net/http"

	"first-go-project/configs"
	"first-go-project/internal/infrastructure/database"
	"first-go-project/internal/infrastructure/database/models"
	"first-go-project/internal/modules/user"
	"first-go-project/internal/router"
	"first-go-project/pkg/response"

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

	// Run migrations
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("✓ Migrations completed")

	// Initialize module handlers
	userHandler := user.NewHandler(db.DB)

	// Setup Gin router
	ginRouter := gin.Default()

	// Health check endpoint
	ginRouter.GET("/ping", func(context *gin.Context) {
		response.Success(context, http.StatusOK, gin.H{
			"service": "first-go-project",
			"status":  "healthy",
		}, "Service is running")
	})

	// Setup routes
	router.SetupUserRoutes(ginRouter, userHandler)

	// Start server
	addr := fmt.Sprintf(":%s", config.App.Port)
	log.Printf("Server starting on %s", addr)
	if err := ginRouter.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
