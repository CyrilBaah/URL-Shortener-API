package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CyrilBaah/URL-Shortener-API/router"
)

func main() {
	r := router.SetupRouter()
	port := "8080"
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
