package handlers

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"backend/internal/api/dto"
	"backend/internal/db"
	"backend/internal/utils"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["id"].(float64))

	return uid == n
}

func validUser(id string, p string) bool {
	var user db.User
	db.GetInstance().First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Token) {
		return false
	}
	return true
}

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user db.User
	db.GetInstance().Find(&user, id)
	if user.Username == "" {
		return c.Status(404).
			JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// CreateUser godoc
// @Summary Создание нового пользователя
// @Description Создает нового пользователя либо по telegram_id либо по email
// @Router /api/user [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(db.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "Invalid request body", "errors": err.Error()})
	}

	user.Token = utils.GenerateRandomPassword()
	fmt.Println("token=", user.Token)
	hash, err := hashPassword(user.Token)
	if err != nil {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	}

	user.Token = hash
	if err := db.GetInstance().Create(&user).Error; err != nil {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}

	newUser := dto.UserAccountDto{
		Id:         user.SerialID,
		Email:      user.Email,
		Handle:     user.Handle,
		TelegramId: user.TelegramID,
		Token:      user.Token,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// TODO: исправить
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Names string `json:"names"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).
			JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	var user db.User

	db.GetInstance().First(&user, id)
	db.GetInstance().Save(&user)

	return c.JSON(
		fiber.Map{"status": "success", "message": "User successfully updated", "data": user},
	)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
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
