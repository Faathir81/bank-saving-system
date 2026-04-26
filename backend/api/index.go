package handler

import (
	"net/http"
	"sync"

	"bank-saving-system/config"
	"bank-saving-system/routes"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	app  *fiber.App
	once sync.Once
)

func setup() {
	// 1. Connect to DB (Neon) - NO AutoMigrate to prevent Serverless Timeout
	config.ConnectDB()

	// 2. Setup Fiber
	app = fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// 3. Register Routes
	routes.SetupRoutes(app)
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(setup)
	adaptor.FiberApp(app).ServeHTTP(w, r)
}
