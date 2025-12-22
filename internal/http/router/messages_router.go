package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igudelj/chat-backend/internal/http/handlers"
)

type ChatRouter struct{}

func (r *ChatRouter) Register(api fiber.Router) {
	api.Get("/chat/messages", handlers.GetMessages)
}
