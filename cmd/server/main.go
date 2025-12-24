package main

import (
	"log"
	"os"

	_ "github.com/igudelj/chat-backend/docs"
	"github.com/igudelj/chat-backend/internal/app"
)

// @title Chat Backend API
// @version 1.0
// @description Backend service for modular chat application
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server, db, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	log.Printf("Starting server on :%s", port)
	log.Fatal(server.Listen(":" + port))
}
