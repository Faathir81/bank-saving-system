package handler

import (
	"net/http"
	"sync"

	"bank-saving-system/config"
	"bank-saving-system/routes"
	"bank-saving-system/utils"
)

var (
	handler http.Handler
	once    sync.Once
)

func setup() {
	// 1. Connect to DB (Neon)
	config.ConnectDB()

	// 2. Setup ServeMux
	mux := http.NewServeMux()

	// 3. Register Routes
	routes.SetupRoutes(mux)

	// 4. Wrap with CORS middleware
	handler = utils.CORS(mux)
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(setup)
	handler.ServeHTTP(w, r)
}
