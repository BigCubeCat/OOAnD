package user

import (
	"backend/internal/db"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUser godoc
// @Summary Получение пользователя по Id
// @Router /api/user/:id [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user db.User
	db.GetInstance().Find(&user, id)
	if user.Handle == "" && user.Email == "" {
		return c.Status(404).
			JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func GetUserByEmail(e string) (*db.User, error) {
	var user db.User
	if err := db.GetInstance().Where(&db.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByTgId(tgId int) (*db.User, error) {
	var user db.User
	if err := db.GetInstance().Where(&db.User{TelegramID: tgId}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserById(id int) (*db.User, error) {
	var user db.User
	if err := db.GetInstance().Where(&db.User{SerialID: id}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	log.Println(user)
	return &user, nil
}
