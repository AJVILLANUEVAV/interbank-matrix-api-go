package main

import (
	"log"
	"os"

	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/adapters"
	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := handler.NewApp(adapters.StatisticsURL(), adapters.JWTSecret())
	log.Fatal(app.Listen(":" + port))
}
