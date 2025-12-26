package message

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/igudelj/chat-backend/internal/entities"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(
	ctx context.Context,
	senderID uuid.UUID,
	receiverID uuid.UUID,
	content string,
) (*entities.Message, error) {
	query := `
		INSERT INTO messages (sender_id, receiver_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	msg := &entities.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
	}

	err := r.db.QueryRowContext(
		ctx,
		query,
		senderID,
		receiverID,
		content,
	).Scan(&msg.ID, &msg.CreatedAt)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (r *postgresRepository) ListBetweenUsers(
	ctx context.Context,
	userA uuid.UUID,
	userB uuid.UUID,
	limit int,
	offset int,
) ([]*entities.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, content, created_at
		FROM messages
		WHERE
			(sender_id = $1 AND receiver_id = $2)
			OR
			(sender_id = $2 AND receiver_id = $1)
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`

	rows, err := r.db.QueryContext(ctx, query, userA, userB, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}(rows)

	var messages []*entities.Message

	for rows.Next() {
		var msg entities.Message
		if err := rows.Scan(
			&msg.ID,
			&msg.SenderID,
			&msg.ReceiverID,
			&msg.Content,
			&msg.CreatedAt,
		); err != nil {
			return nil, err
		}

		messages = append(messages, &msg)
	}

	return messages, rows.Err()
}
