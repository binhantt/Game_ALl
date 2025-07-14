package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/game_all?charset=utf8mb4&parseTime=True&loc=Local"
	
	// Try to connect to the database with retries
	var db *gorm.DB
	var err error
	maxRetries := 3

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			log.Println("✅ Successfully connected to the database")
			return
		}
		
		log.Printf("⚠️  Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(2 * time.Second) // Wait before retrying
		}
	}

	// If we get here, we couldn't connect to the database
	log.Println("❌ WARNING: Could not connect to the database. The application will start without database connection.")
	log.Println("❗ To fix this, please ensure:")
	log.Println("1. MySQL server is running")
	log.Println("2. Database 'game_all' exists")
	log.Println("3. User 'root' has the correct password")
	log.Println("4. MySQL is accessible at 127.0.0.1:3306")
}
