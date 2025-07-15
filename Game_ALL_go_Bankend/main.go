package main

import (
	"Game_ALL_go_bankend/config"
	"Game_ALL_go_bankend/pkg/migrate"
	"Game_ALL_go_bankend/routes"
	"log"
)

func main() {
	// Initialize database connection (this will continue even if DB connection fails)
	config.ConnectDB()
	// Run database migrations
	if config.DB != nil {
		if err := migrate.AutoMigrate(config.DB); err != nil {
			log.Printf("⚠️  Failed to run database migrations: %v", err)
		} else {
			log.Println("✅ Database migrations completed successfully")
		}
	}

	// Create and start the server
	server := config.NewServer("8081")

	// Setup all API routes
	routes.SetupRoutes(server)

	// Start the server
	log.Println("Starting server on port 8081...")
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
