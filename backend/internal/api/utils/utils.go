package api

import "github.com/gofiber/fiber/v2"

func CreatePrettyError(c *fiber.Ctx, status int, message string, data error) error {
	return c.Status(status).
		JSON(fiber.Map{"status": "error", "message": message, "data": data})
}
