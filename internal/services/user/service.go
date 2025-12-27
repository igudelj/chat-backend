package user

import (
	"context"

	"github.com/igudelj/chat-backend/internal/entities"
)

type Service interface {
	Search(ctx context.Context, field entities.UserSearchField, value string) (*entities.User, error)
	Create(ctx context.Context, user *entities.User, password string) error
}
