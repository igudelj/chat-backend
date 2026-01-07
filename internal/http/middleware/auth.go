package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func KeycloakJWT(jwksURL string) fiber.Handler {
	// Fetch JWKS manually
	resp, err := http.Get(jwksURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var jwks map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		panic(err)
	}

	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return fiber.ErrUnauthorized
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		tokenStr := parts[1]

		// Parse token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if t.Method.Alg() != jwt.SigningMethodRS256.Alg() {
				return nil, fmt.Errorf("invalid alg")
			}

			kid, ok := t.Header["kid"].(string)
			if !ok {
				return nil, fmt.Errorf("missing kid")
			}

			// Look up key in JWKS
			keys := jwks["keys"].([]interface{})
			for _, k := range keys {
				key := k.(map[string]interface{})
				if key["kid"] == kid {
					certPEM := "-----BEGIN PUBLIC KEY-----\n" + key["x5c"].([]interface{})[0].(string) + "\n-----END PUBLIC KEY-----"
					pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(certPEM))
					if err != nil {
						return nil, err
					}
					return pubKey, nil
				}
			}
			return nil, fmt.Errorf("key not found")
		})
		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.ErrUnauthorized
		}

		c.Locals("token", token)
		c.Locals("claims", claims)

		return c.Next()
	}
}
