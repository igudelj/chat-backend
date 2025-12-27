package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igudelj/chat-backend/internal/entities"
	userservice "github.com/igudelj/chat-backend/internal/services/user"
)

type Handler struct {
	service userservice.Service
}

func New(service userservice.Service) *Handler {
	return &Handler{service: service}
}

// Search godoc
// @Summary Search users
// @Description Search users by a single field (id, email, username)
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
func (h *Handler) Search(c *fiber.Ctx) error {
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
// @Tags user
// @Accept json
// @Produce json
// @Param request body createUserRequest true "Create user payload"
// @Success 201 {object} entities.User
// @Failure 400
// @Failure 500
// @Router /users [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	user := &entities.User{
		Username: req.Username,
		Email:    req.Email,
	}

	err := h.service.Create(c.Context(), user, req.Password)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
