//go:build ignore

package main

import (
	"Game_ALL_go_bankend/config"
	"Game_ALL_go_bankend/pkg/migrate"
	"log"
)

func main() {
	// Initialize database connection
	config.ConnectDB()

	// Run migrations
	if err := migrate.AutoMigrate(config.DB); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ All migrations completed successfully")
}
