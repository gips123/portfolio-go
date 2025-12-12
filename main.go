package main

import (
	"best-portfolio-go/config"
	"best-portfolio-go/database"
	"best-portfolio-go/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Log CORS configuration
	log.Printf("CORS Allowed Origins: %v", cfg.CORS.AllowOrigins)

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup routes
	router := routes.SetupRoutes(cfg)

	// Get port from config
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
