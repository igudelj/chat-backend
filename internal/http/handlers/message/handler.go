package message

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igudelj/chat-backend/internal/entities"
	messageservice "github.com/igudelj/chat-backend/internal/services/message"
)

type Handler struct {
	service messageservice.Service
}

func New(service messageservice.Service) *Handler {
	return &Handler{service: service}
}

// GetMessages godoc
// @Summary Get chat messages
// @Description Returns messages between two users
// @Tags chat
// @Accept json
// @Produce json
// @Param user_a query string true "User A ID"
// @Param user_b query string true "User B ID"
// @Param limit query int false "Limit" default(50)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} entities.Message
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /chat/messages [get]
func (h *Handler) GetMessages(c *fiber.Ctx) error {
	userA, err := uuid.Parse(c.Query("user_a"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_a")
	}

	userB, err := uuid.Parse(c.Query("user_b"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_b")
	}

	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	messages, err := h.service.GetConversation(
		c.Context(),
		userA,
		userB,
		limit,
		offset,
	)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(messages)
}

type sendMessageRequest struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
}

// SendMessage godoc
// @Summary Send a message
// @Description Sends a text message
// @Tags chat
// @Accept json
// @Produce json
// @Param request body sendMessageRequest true "Message payload"
// @Success 201 {object} entities.Message
// @Failure 400
// @Failure 500
// @Router /chat/messages [post]
func (h *Handler) SendMessage(c *fiber.Ctx) error {
	var req sendMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	senderID, err := uuid.Parse(req.SenderID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid sender_id")
	}

	receiverID, err := uuid.Parse(req.ReceiverID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid receiver_id")
	}

	if err, _ := h.service.Send(c.Context(), senderID, receiverID, req.Content); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(&entities.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    req.Content,
	})
}
