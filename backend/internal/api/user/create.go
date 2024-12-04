package user

import (
	"backend/internal/api/dto"
	"backend/internal/db"
	"backend/internal/utils"

	apiUtils "backend/internal/api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
// @Summary Создание нового пользователя
// @Description Создает нового пользователя либо по telegram_id либо по email
// @Router /api/user [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(db.User)
	if err := c.BodyParser(user); err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Review your input", err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "Invalid request body", "errors": err.Error()})
	}

	user.Token = utils.GenerateRandomPassword()
	hash, err := HashPassword(user.Token)
	if err != nil {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	}

	user.Token = hash
	if err := db.GetInstance().Create(&user).Error; err != nil {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}
	t, err := createToken(user.SerialID)
	if err != nil {
		return apiUtils.CreatePrettyError(c, fiber.StatusInternalServerError, err.Error(), err)
	}

	newUser := dto.UserAccountDto{
		Id:         user.SerialID,
		Email:      user.Email,
		Handle:     user.Handle,
		TelegramId: user.TelegramID,
		Token:      t,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}
