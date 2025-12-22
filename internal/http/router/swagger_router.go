package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type SwaggerRouter struct{}

func (r *SwaggerRouter) Register(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
