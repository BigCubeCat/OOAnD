package user

import (
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !ValidToken(token, id) {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	var user db.User

	db.GetInstance().First(&user, id)

	db.GetInstance().Delete(&user)
	return c.JSON(
		fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil},
	)
}
