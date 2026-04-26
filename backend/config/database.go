package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	
	// Check for DATABASE_URL (Standard for Neon/Render/Railway)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback to local individual params
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASS", ""),
			getEnv("DB_NAME", "bank_saving_db"),
			getEnv("DB_PORT", "5432"),
		)
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		WithoutQuotingCheck:  true, // Disable pg_catalog queries to speed up Vercel cold start
		PreferSimpleProtocol: true, // Disables implicit prepared statement usage
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Println("Failed to connect to database:", err)
		// Don't log.Fatal, let the handler panic or return 500 later, otherwise Vercel hard crashes.
	}

	fmt.Println("Connected to Database!")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
