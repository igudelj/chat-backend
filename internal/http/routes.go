package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igudelj/chat-backend/internal/http/router"
)

func RegisterRoutes(app *fiber.App) {
	// Non-versioned routes
	(&router.SwaggerRouter{}).Register(app)

	// rest of the routes
	apiV1 := app.Group("/api/v1")
	routers := []router.Router{
		&router.HealthRouter{},
		&router.ChatRouter{},
	}

	for _, r := range routers {
		r.Register(apiV1)
	}
}
