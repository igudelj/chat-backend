package handlers

import "github.com/gofiber/fiber/v2"

// GetHealth godoc
// @Summary Get service health check status
// @Description Returns server status check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} {"status": "ok"}
// @Failure 500
// @Router /health [get]
func GetHealth(c *fiber.Ctx) error {
	return c.JSON([]fiber.Map{
		{"status": "ok"},
	})
}
