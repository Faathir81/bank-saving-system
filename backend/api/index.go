package handler

import (
	"net/http"
	"strings"
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

	// 4. Wrap with CORS and Panic Recovery middleware
	handler = utils.Middleware(mux)
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(setup)

	// Vercel Rewrite Fix: If r.URL.Path is /api/index.go, we shouldn't serve it.
	// We need to restore the actual requested path from headers if Vercel mangled it.
	if forwardedURL := r.Header.Get("x-now-route-matches"); forwardedURL != "" {
		// Vercel sometimes puts internal routing info here, but usually r.URL.Path is fine.
		// However, just in case Vercel stripped "/api", we can add a quick check.
	}
	
	// If the path is literally /api/index.go, it means the rewrite destination leaked into Path.
	// This causes 404s. We should use the original URL if available.
	if strings.Contains(r.URL.Path, "index.go") {
		originalPath := r.Header.Get("x-invoke-path")
		if originalPath != "" {
			r.URL.Path = originalPath
		}
	}

	handler.ServeHTTP(w, r)
}
