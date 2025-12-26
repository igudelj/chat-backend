package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/igudelj/chat-backend/internal/entities"
)

type Repository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}
