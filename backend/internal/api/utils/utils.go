package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func createPrettyMessage(
	c *fiber.Ctx,
	statusCode int,
	status string,
	message string,
	data interface{},
) error {
	return c.Status(statusCode).
		JSON(fiber.Map{"status": status, "message": message, "data": data})
}

func CreatePrettyError(c *fiber.Ctx, status int, message string, err error) error {
	log.Error(err.Error())
	return createPrettyMessage(c, status, "error", message, nil)
}

func CreatePrettySuccess(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{"status": "success", "data": data})
}

func ParseId(c *fiber.Ctx) (int, error) {
	idParam := c.Params("id")
	return strconv.Atoi(idParam)
}
