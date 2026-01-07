package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/igudelj/chat-backend/internal/entities"
	userservice "github.com/igudelj/chat-backend/internal/services/user"
)

type UserHandler struct {
	service userservice.Service
}

func NewUserHandler(service userservice.Service) *UserHandler {
	return &UserHandler{service: service}
}

// Search godoc
// @Summary Search users
// @Description Search users by a single field (id, email, username)
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param id query string false "User ID (UUID)"
// @Param email query string false "User email"
// @Param username query string false "Username"
// @Success 200 {object} entities.User
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /users [get]
func (h *UserHandler) Search(c *fiber.Ctx) error {
	query := c.Queries()

	if len(query) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "at least one search parameter is required")
	}

	if len(query) > 1 {
		return fiber.NewError(fiber.StatusBadRequest, "only one search parameter is allowed")
	}

	var (
		field entities.UserSearchField
		value string
	)

	for k, v := range query {
		parsed, ok := entities.ParseUserSearchField(k)
		if !ok {
			return fiber.NewError(
				fiber.StatusBadRequest,
				"unsupported search parameter: "+k,
			)
		}
		field = parsed
		value = v
	}

	user, err := h.service.Search(c.Context(), field, value)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if user == nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser godoc
// @Summary Create user
// @Description Creates a new user and returns the created object.
// @Security BearerAuth
// @Tags user
// @Accept json
// @Produce json
// @Param request body createUserRequest true "Create user payload"
// @Success 201 {object} entities.User
// @Failure 400
// @Failure 500
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	if claims == nil {
		return fiber.ErrUnauthorized
	}

	// Ensure the user exists in local DB (or create if missing)
	user, err := h.service.EnsureCurrentUser(c.Context(), claims)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to ensure user")
	}

	// Return DB info + Keycloak claims
	err = c.JSON(fiber.Map{
		"id":          user.ID,
		"keycloak_id": user.KeycloakID,
		"username":    claims["preferred_username"],
		"email":       claims["email"],
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
