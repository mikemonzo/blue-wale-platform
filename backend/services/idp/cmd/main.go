package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/config"
)

func main() {
	cfg := config.NewConfig()

	setupRoutes()

	log.Printf("Starting IDP service on :%d", cfg.Server.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setupRoutes() {
	http.HandleFunc("/health", healthHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
