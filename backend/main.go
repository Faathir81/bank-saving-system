package main

import (
	"log"

	"bank-saving-system/config"
	"bank-saving-system/models"
	"bank-saving-system/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to Database
	config.ConnectDB()

	// Auto Migration
	err := config.DB.AutoMigrate(
		&models.Customer{},
		&models.DepositoType{},
		&models.Account{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Setup Routes
	routes.SetupRoutes(app)

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
