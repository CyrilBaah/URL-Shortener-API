package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/CyrilBaah/URL-Shortener-API/handler"
)

// HealthCheckHandler responds with "OK" for readiness/liveness probes
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🟢 Health check endpoint accessed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

// SetupRouter initializes the API routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true) // Ensure strict matching

	// Middleware to log every request
	r.Use(loggingMiddleware)

	// Register the health check route FIRST to avoid conflicts
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	log.Println("✅ Route registered: GET /health")

	// Register API routes
	r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	log.Println("✅ Route registered: POST /shorten")

	r.HandleFunc("/{shortURL}", handler.ResolveURL).Methods("GET")
	log.Println("✅ Route registered: GET /{shortURL}")

	// Log all registered routes
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			log.Println("📌 Registered route:", path)
		}
		return nil
	})
	if err != nil {
		log.Println("❌ Error listing routes:", err)
	}

	return r
}

// Middleware to log incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🔍 Incoming request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
