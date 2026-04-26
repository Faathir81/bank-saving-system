package handler

import (
	"net/http"
	"bank-saving-system/config"
	"bank-saving-system/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/adaptor/v2"
)

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. Connect to DB (Neon)
	config.ConnectDB()
	
	// 2. Setup Fiber
	app := fiber.New()
	app.Use(cors.New())
	
	// 3. Register Routes
	routes.SetupRoutes(app)
	
	// 4. Adapt Fiber to Standard Go Http Handler
	adaptor.FiberApp(app).ServeHTTP(w, r)
}
