package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/igudelj/chat-backend/internal/entities"
	"github.com/igudelj/chat-backend/internal/repositories/user"
)

type service struct {
	userRepo user.Repository
}

func New(userRepo user.Repository) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *service) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}
