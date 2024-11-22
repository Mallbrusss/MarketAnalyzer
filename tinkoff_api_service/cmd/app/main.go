package main

import (
	"log"
	"net/http"
	"tinkoff-api/config"
	"tinkoff-api/internal/handlers"
	"tinkoff-api/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	service := services.NewTinkoffService(cfg)
	handler := handlers.NewHandler(service)

	http.HandleFunc("/api/v1/closePrices", handler.GetClosePricesHandler)

	log.Printf("Server is running on port %s...", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
