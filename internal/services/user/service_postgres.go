package user

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
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

func (s *service) Search(ctx context.Context, field entities.UserSearchField, value string) (*entities.User, error) {
	return s.userRepo.GetByField(ctx, field, value)
}

func (s *service) Create(ctx context.Context, user *entities.User, password string) error {
	return s.userRepo.Create(ctx, user)
}

func (s *service) EnsureCurrentUser(ctx context.Context, claims jwt.MapClaims) (*entities.User, error) {
	return s.userRepo.EnsureFromClaims(ctx, claims)
}
