package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/CyrilBaah/URL-Shortener-API/handler"
)

// Middleware to track requests in Prometheus
func metricsMiddleware(httpRequestsTotal *prometheus.CounterVec) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
			next.ServeHTTP(w, r)
		})
	}
}

// Health check handler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🟢 Health check endpoint accessed")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// SetupRouter initializes the API routes (Fixed: Adds `/metrics`)
func SetupRouter(httpRequestsTotal *prometheus.CounterVec) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Middleware for logging & metrics
	r.Use(metricsMiddleware(httpRequestsTotal))

	// Register the Prometheus `/metrics` route
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	log.Println("✅ Route registered: GET /metrics")

	// Register health check
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	log.Println("✅ Route registered: GET /health")

	// Register API routes
	r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	log.Println("✅ Route registered: POST /shorten")

	r.HandleFunc("/{shortURL}", handler.ResolveURL).Methods("GET")
	log.Println("✅ Route registered: GET /{shortURL}")

	return r
}
