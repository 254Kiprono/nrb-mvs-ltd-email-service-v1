package main

import (
	"email-service/config"
	"email-service/database"
	"email-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.InitializeDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set up routes
	routes.SetupEmailRoutes(r, db)

	// Start server
	port := "9015"
	log.Printf("Starting server on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
