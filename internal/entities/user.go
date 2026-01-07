package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`          // internal app user ID
	KeycloakID uuid.UUID `json:"keycloak_id"` // Keycloak sub
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
