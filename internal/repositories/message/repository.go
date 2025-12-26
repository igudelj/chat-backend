package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/igudelj/chat-backend/internal/entities"
)

type Repository interface {
	Create(
		ctx context.Context,
		senderID uuid.UUID,
		receiverID uuid.UUID,
		content string,
	) (*entities.Message, error)
	ListBetweenUsers(
		ctx context.Context,
		userA uuid.UUID,
		userB uuid.UUID,
		limit int,
		offset int,
	) ([]*entities.Message, error)
}
