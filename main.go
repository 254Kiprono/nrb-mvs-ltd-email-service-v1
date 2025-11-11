package main

import (
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

	// Set up email routes
	routes.SetupEmailRoutes(r) 

	// Start server
	port := "9014" // Default port
	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
