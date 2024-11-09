package handlers

import (
	"backend/internal/config"
	"backend/internal/db"
	"errors"
	"fmt"
	"net/mail"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*db.User, error) {
	var user db.User
	if err := db.GetInstance().Where(&db.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByTgId(tgId int) (*db.User, error) {
	var user db.User
	if err := db.GetInstance().Where(&db.User{TelegramID: tgId}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Token    string `json:"token"`
	}
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Error on login request", "errors": err.Error()})
	}
	fmt.Println(input)
	userModel, err := new(db.User), *new(error)
	if validEmail(input.Identity) {
		userModel, err = getUserByEmail(input.Identity)
	} else {
		tgId, err := strconv.Atoi(input.Identity)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"status": "error", "message": "Invalid Telegram ID", "errors": err.Error()})
		}
		userModel, err = getUserByTgId(tgId)
	}
	fmt.Println("userModel = ", userModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": "Internal Server Error", "data": err})
	}
	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": err})
	}
	if !CheckPasswordHash(input.Token, userModel.Token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error", "message": "Invalid token", "data": nil,
		})
	}
	token := jwt.New(jwt.SigningMethodES256)
	fmt.Println(token)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userModel.SerialID
	claims["email"] = userModel.Email
	claims["tg_id"] = userModel.TelegramID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println("claims = ", claims)

	t, err := token.SignedString([]byte(config.GetJwtSecret()))
	fmt.Println("t=", t)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(), "data": nil,
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
