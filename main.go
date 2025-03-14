package main

import (
	"log"
	"net/http"

	"github.com/CyrilBaah/URL-Shortener-API/router"
	"github.com/prometheus/client_golang/prometheus"
)

// Define Prometheus metrics (GLOBAL)
var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "endpoint"},
)

func main() {
	// Register Prometheus metrics (ONLY ONCE)
	prometheus.MustRegister(httpRequestsTotal)

	// Pass the metric to SetupRouter
	r := router.SetupRouter(httpRequestsTotal)

	port := "8080"
	log.Println("🚀 Server running on port", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
