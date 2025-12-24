package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/igudelj/chat-backend/internal/config"
	"github.com/igudelj/chat-backend/internal/database"
	"github.com/igudelj/chat-backend/internal/http"
)

func New() (*fiber.App, *sql.DB, error) {
	app := fiber.New()

	cfg := config.LoadDatabaseConfig()

	db, err := database.PostgresInstance(cfg)
	if err != nil {
		return nil, nil, err
	}

	http.RegisterRoutes(app)

	return app, db, nil
}
