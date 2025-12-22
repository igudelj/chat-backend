package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/igudelj/chat-backend/docs"
	"github.com/igudelj/chat-backend/internal/http"
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
	app := fiber.New()

	http.RegisterRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "running"})
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
