package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igudelj/chat-backend/internal/http"
)

func New() *fiber.App {
	app := fiber.New()

	http.RegisterRoutes(app)

	return app
}
