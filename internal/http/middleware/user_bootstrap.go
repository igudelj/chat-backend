package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/igudelj/chat-backend/internal/repositories/user"
)

func UserBootstrap(userRepo user.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.Locals("token").(*jwt.Token)
		if !ok {
			return fiber.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.ErrUnauthorized
		}

		userFromClaims, err := userRepo.EnsureFromClaims(c.Context(), claims)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		// Store domain user for downstream handlers
		c.Locals("user", userFromClaims)
		c.Locals("userID", userFromClaims.ID)

		return c.Next()
	}
}
