package main

import (
	"artifact-store/internal/api"
	"net/http"
)

func main() {
	server := api.NewServer()
	router := http.NewServeMux()
	handler := api.HandlerFromMux(server, router)
	service := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:8080",
	}
	service.ListenAndServe()
}
