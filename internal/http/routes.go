package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/igudelj/chat-backend/internal/http/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Get("/swagger/*", swagger.HandlerDefault)
	api.Get("/chat/messages", handlers.GetMessages)
}
