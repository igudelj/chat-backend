package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	userservice "github.com/igudelj/chat-backend/internal/services/user"
)

type Handler struct {
	service userservice.Service
}

func New(service userservice.Service) *Handler {
	return &Handler{service: service}
}

// GetByID godoc
// @Summary Get user by ID
// @Description Returns a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} entities.User
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /users/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	user, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if user == nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

// GetByEmail godoc
// @Summary Get user by email
// @Description Returns a user by email
// @Tags users
// @Accept json
// @Produce json
// @Param email query string true "User email"
// @Success 200 {object} entities.User
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /users [get]
func (h *Handler) GetByEmail(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}

	user, err := h.service.GetByEmail(c.Context(), email)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if user == nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}
