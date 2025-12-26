package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/igudelj/chat-backend/internal/entities"
	"github.com/igudelj/chat-backend/internal/repositories/message"
)

type service struct {
	messageRepo message.Repository
}

func New(messageRepo message.Repository) Service {
	return &service{
		messageRepo: messageRepo,
	}
}

func (s *service) Send(
	ctx context.Context,
	senderID uuid.UUID,
	receiverID uuid.UUID,
	content string,
) (*entities.Message, error) {
	return s.messageRepo.Create(ctx, senderID, receiverID, content)
}

func (s *service) GetConversation(
	ctx context.Context,
	userA uuid.UUID,
	userB uuid.UUID,
	limit int,
	offset int,
) ([]*entities.Message, error) {
	// Guardrails (important)
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	return s.messageRepo.ListBetweenUsers(ctx, userA, userB, limit, offset)
}
