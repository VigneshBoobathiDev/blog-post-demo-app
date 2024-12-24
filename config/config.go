package config

import (
	"blogpost/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// LoadEnv loads environment variables from a .env file.
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// ConnectDB connects to the database and returns a *gorm.DB instance.
func ConnectDB() *gorm.DB {
	// Retrieve environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	dbname := os.Getenv("DB_NAME")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbname)

	// Open a database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connected successfully!")

	// Auto-migrate models
	if err := db.AutoMigrate(&models.Article{}, &models.Comment{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	return db
}
