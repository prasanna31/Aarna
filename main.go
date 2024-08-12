package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prasanna31/Aarna/config"
	"github.com/prasanna31/Aarna/middleware"
	"github.com/prasanna31/Aarna/routes"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Create a new Gin router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(middleware.CORS())

	// Apply JWT authentication middleware to protected routes
	router.Use(middleware.AuthMiddleware())

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
