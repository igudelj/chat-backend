package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
		INSERT INTO users (keycloak_id)
		VALUES ($1)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		user.KeycloakID,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *postgresRepository) GetByField(
	ctx context.Context,
	field entities.UserSearchField,
	value any,
) (*entities.User, error) {
	query := fmt.Sprintf(`
		SELECT id, keycloak_id, created_at, updated_at
		FROM users
		WHERE %s = $1
	`, field)

	row := r.db.QueryRowContext(ctx, query, value)

	user := &entities.User{}
	err := row.Scan(
		&user.ID,
		&user.KeycloakID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *postgresRepository) EnsureFromClaims(
	ctx context.Context,
	claims jwt.MapClaims,
) (*entities.User, error) {

	sub := claims["sub"].(string)
	keycloakID, err := uuid.Parse(sub)
	if err != nil {
		return nil, err
	}

	user, err := r.GetByField(ctx, entities.UserKeycloakFieldID, &keycloakID)
	if err == nil {
		return user, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	user = &entities.User{
		ID:         uuid.New(),
		KeycloakID: keycloakID,
	}

	if err := r.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
