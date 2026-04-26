package main

import (
	"log"
	"net/http"

	"bank-saving-system/config"
	"bank-saving-system/models"
	"bank-saving-system/routes"
	"bank-saving-system/utils"
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

	mux := http.NewServeMux()

	// Setup Routes
	routes.SetupRoutes(mux)

	// Wrap mux with CORS middleware
	handler := utils.CORS(mux)

	// Start Server
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
