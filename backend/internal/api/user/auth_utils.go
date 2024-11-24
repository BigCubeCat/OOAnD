package user

import (
	"backend/internal/config"
	"backend/internal/db"
	"log"
	"net/mail"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	apiUtils "backend/internal/api/utils"
)

// Login godoc
// @Summary Аутентификация
// @Description Вход через Телеграм или через почту. Создает пользователя, если его нет
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	tgId := c.FormValue("telegram_id")
	email := c.FormValue("email")
	password := c.FormValue("password")

	userModel, err := new(db.User), *new(error)
	if email == "test@mail.ru" || password == "" {
		// УДАЛИТЬ НА ПРОДЕ
	} else if ValidEmail(email) {
		userModel, err = GetUserByEmail(email)
		if err != nil {
			// регистрация
		}
	} else if ValidTelegramId(tgId) {
		// ошибка была проверена в ValidTelegramId
		tgIdInt, _ := strconv.Atoi(tgId)
		userModel, err = GetUserByTgId(tgIdInt)
		if err != nil {
			// регистрация
		}
	}
	if userModel == nil {
		return apiUtils.CreatePrettyError(
			c,
			fiber.StatusUnauthorized,
			"Invalid identity or password",
			err,
		)
	}
	log.Println("SerialID=", userModel.SerialID)
	t, err := createToken(userModel.SerialID)
	if err != nil {
		return apiUtils.CreatePrettyError(c, fiber.StatusInternalServerError, err.Error(), err)
	}
	return apiUtils.CreatePrettySuccess(c, t)
}

func createToken(id int) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJwtSecret()))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func getIdFromToken(t *jwt.Token) int {
	claims := t.Claims.(jwt.MapClaims)
	return int(claims["id"].(float64))
}

func GetCurrentUserId(c *fiber.Ctx) int {
	user, err := GetCurrentUser(c)
	if err != nil {
		return 0
	}
	return user.SerialID
}

func GetCurrentUser(c *fiber.Ctx) (*db.User, error) {
	token := c.Locals("user").(*jwt.Token)
	log.Println("token=", token)
	tokenString := getIdFromToken(token)
	log.Println("tokenString=", token)
	return GetUserById(tokenString)
}

func ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	return getIdFromToken(t) == n
}

func ValidUser(id string) bool {
	var user db.User
	db.GetInstance().First(&user, id)
	return user.SerialID > 0
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidTelegramId(telegramId string) bool {
	if telegramId == "" {
		return false
	}
	_, err := strconv.Atoi(telegramId)
	return err == nil
}
