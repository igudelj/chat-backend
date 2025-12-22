package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igudelj/chat-backend/internal/http/handlers"
)

type HealthRouter struct{}

func (r *HealthRouter) Register(api fiber.Router) {
	api.Get("/health", handlers.GetHealth)
}
