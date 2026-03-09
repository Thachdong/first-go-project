package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"first-go-project/configs"
	"first-go-project/internal/infrastructure/database"
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

	// Run versioned migrations
	log.Println("Running database migrations...")
	migrationsPath, err := filepath.Abs("internal/infrastructure/database/migrations")
	if err != nil {
		log.Fatalf("Failed to resolve migrations path: %v", err)
	}
	if err := db.RunMigrations(migrationsPath); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

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
	router.SetupUserRoutes(ginRouter, user.NewHandler(db.DB))

	// Start server
	addr := fmt.Sprintf(":%s", config.App.Port)
	log.Printf("Server starting on %s", addr)
	if err := ginRouter.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
