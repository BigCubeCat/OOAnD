package user

import (
	"backend/internal/db"
	"errors"

	apiUtils "backend/internal/api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// UpdateUser godoc
// @Summary Обновление пользователя
// @Description Обновляет поля пользовател
// @Router /api/user/:id [put]
func UpdateUser(c *fiber.Ctx) error {
	var newData db.User
	if err := c.BodyParser(&newData); err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Review your input", err)
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !ValidToken(token, id) {
		return apiUtils.CreatePrettyError(
			c,
			500,
			"Invalid token id",
			errors.New("Invalid token id"),
		)
	}

	var user db.User
	db.GetInstance().First(&user, id)

	user.Email = newData.Email
	user.Handle = newData.Handle
	user.TelegramID = newData.TelegramID

	db.GetInstance().Save(&user)
	return apiUtils.CreatePrettySuccess(c, user)
}
