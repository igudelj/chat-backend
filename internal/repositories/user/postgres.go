package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/igudelj/chat-backend/internal/entities"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, user *entities.User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *postgresRepository) GetByField(
	ctx context.Context,
	field entities.UserSearchField,
	value any,
) (*entities.User, error) {
	query := fmt.Sprintf(`
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE %s = $1
	`, field)

	row := r.db.QueryRowContext(ctx, query, value)

	user := &entities.User{}
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
