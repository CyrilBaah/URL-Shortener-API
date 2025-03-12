package router

import (
	"github.com/gorilla/mux"
	"github.com/CyrilBaah/URL-Shortener-API/handler"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", handler.ResolveURL).Methods("GET")
	return r
}
