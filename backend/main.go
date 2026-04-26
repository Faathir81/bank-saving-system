package main

import (
	"log"
	"net/http"
	"os"

	"bank-saving-system/config"
	"bank-saving-system/models"
	"bank-saving-system/routes"
	"bank-saving-system/utils"
)

func main() {
	// Connect to Database
	config.ConnectDB()

	// Only AutoMigrate if NOT in Vercel (Vercel sets VERCEL=1)
	if os.Getenv("VERCEL") == "" {
		err := config.DB.AutoMigrate(
			&models.Customer{},
			&models.DepositoType{},
			&models.Account{},
			&models.Transaction{},
		)
		if err != nil {
			log.Fatal("Migration failed:", err)
		}
	}

	mux := http.NewServeMux()

	// Setup Routes
	routes.SetupRoutes(mux)

	// Wrap mux with CORS and Panic Recovery middleware
	handler := utils.Middleware(mux)

	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start Server
	log.Println("Server listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
