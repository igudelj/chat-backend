package router

import (
	"github.com/gofiber/fiber/v2"
	messagehandler "github.com/igudelj/chat-backend/internal/http/handlers/message"
)

type MessagesRouter struct {
	messageHandler *messagehandler.Handler
}

func NewMessagesRouter(handler *messagehandler.Handler) *MessagesRouter {
	return &MessagesRouter{
		messageHandler: handler,
	}
}

func (r *MessagesRouter) Register(api fiber.Router) {
	api.Get("/chat/messages", r.messageHandler.GetMessages)
	api.Post("/chat/messages", r.messageHandler.SendMessage)
}
