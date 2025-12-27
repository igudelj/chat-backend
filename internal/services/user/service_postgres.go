package user

import (
	"context"

	"github.com/igudelj/chat-backend/internal/entities"
	"github.com/igudelj/chat-backend/internal/repositories/user"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hash)

	return s.userRepo.Create(ctx, user)
}
