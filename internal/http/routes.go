package http

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/igudelj/chat-backend/internal/http/handlers"
	"github.com/igudelj/chat-backend/internal/http/middleware"

	"github.com/igudelj/chat-backend/internal/http/router"
	messagerepo "github.com/igudelj/chat-backend/internal/repositories/message"
	userrepo "github.com/igudelj/chat-backend/internal/repositories/user"
	messageservice "github.com/igudelj/chat-backend/internal/services/message"
	userservice "github.com/igudelj/chat-backend/internal/services/user"
)

func RegisterRoutes(app *fiber.App, db *sql.DB) {
	// Non-versioned routes
	(&router.SwaggerRouter{}).Register(app)

	// repositories
	messageRepo := messagerepo.NewPostgresRepository(db)
	userRepo := userrepo.NewPostgresRepository(db)

	// services
	messageService := messageservice.New(messageRepo)
	userService := userservice.New(userRepo)

	// handlers
	messageHandler := handlers.NewMessageHandler(messageService)
	userHandler := handlers.NewUserHandler(userService)

	// attach auth middleware
	keycloakUrl := os.Getenv("KC_HOSTNAME_URL")
	keycloakRealm := os.Getenv("KC_IMPORT_REALM")
	jwksURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/certs", keycloakUrl, keycloakRealm)
	print(" keycloak url: " + keycloakUrl)
	print(" keycloak realm: " + keycloakRealm)
	print(" keycloak jwks url: " + jwksURL)

	app.Use(middleware.KeycloakJWT(jwksURL))
	app.Use(middleware.UserBootstrap(userRepo))

	// rest of the routes
	apiV1 := app.Group("/api/v1")
	routers := []router.Router{
		&router.HealthRouter{},
		router.NewMessagesRouter(messageHandler),
		router.NewUsersRouter(userHandler),
	}

	for _, r := range routers {
		r.Register(apiV1)
	}
}
