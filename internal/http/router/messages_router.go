package router

import (
	"github.com/gofiber/fiber/v2"
	messagehandler "github.com/igudelj/chat-backend/internal/http/handlers"
)

type MessagesRouter struct {
	messageHandler *messagehandler.MessageHandler
}

func NewMessagesRouter(handler *messagehandler.MessageHandler) *MessagesRouter {
	return &MessagesRouter{
		messageHandler: handler,
	}
}

func (r *MessagesRouter) Register(api fiber.Router) {
	api.Get("/chat/messages", r.messageHandler.GetMessages)
	api.Post("/chat/messages", r.messageHandler.SendMessage)
}
