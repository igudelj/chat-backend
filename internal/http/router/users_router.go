package router

import (
	"github.com/gofiber/fiber/v2"
	usershandler "github.com/igudelj/chat-backend/internal/http/handlers/user"
)

type UsersRouter struct {
	usersHandler *usershandler.Handler
}

func NewUsersRouter(handler *usershandler.Handler) *UsersRouter {
	return &UsersRouter{usersHandler: handler}
}

func (r *UsersRouter) Register(api fiber.Router) {
	api.Get("/users", r.usersHandler.Search)
	api.Post("/users", r.usersHandler.CreateUser)
}
