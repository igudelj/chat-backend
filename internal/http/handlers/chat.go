package handlers

import "github.com/gofiber/fiber/v2"

// GetMessages godoc
// @Summary Get chat messages
// @Description Returns recent chat messages
// @Tags chat
// @Accept json
// @Produce json
// @Success 200 {array} map[string]string
// @Failure 500 {object} map[string]string
// @Router /chat/messages [get]
func GetMessages(c *fiber.Ctx) error {
	return c.JSON([]fiber.Map{
		{"message": "hello"},
	})
}
