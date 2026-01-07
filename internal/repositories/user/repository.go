package user

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/igudelj/chat-backend/internal/entities"
)

type Repository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByField(
		ctx context.Context,
		field entities.UserSearchField,
		value any,
	) (*entities.User, error)
	EnsureFromClaims(
		ctx context.Context,
		claims jwt.MapClaims,
	) (*entities.User, error)
}
