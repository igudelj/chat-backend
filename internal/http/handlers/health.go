package handlers

import "github.com/gofiber/fiber/v2"

type HealthResponse struct {
	Status string `json:"status"`
}

// GetHealth godoc
// @Summary Get service health check status
// @Description Returns server status check
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Failure 500
// @Router /health [get]
func GetHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(HealthResponse{
		Status: "ok",
	})
}
