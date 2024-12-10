package user

import (
	"backend/internal/config"
	"backend/internal/db"
	"errors"
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
	tgId := c.FormValue("id")
	username := c.FormValue("username")
	password := c.FormValue("password")
	photo_url := c.FormValue("photo_url")
	first_name := c.FormValue("first_name")
	last_name := c.FormValue("last_name")

	userModel, err := new(db.User), *new(error)
	if ValidTelegramId(tgId) {
		// Авторизация через телеграм
		tgIdInt, _ := strconv.Atoi(tgId) // ошибка проверяется в ValidTelegramId
		userModel, err = GetUserByTgId(tgIdInt)
		if err != nil {
			// регистрация
			return CreateUser(c, tgIdInt, last_name, first_name, username, photo_url)
		} else {
			// логин
			if err != nil {
				return apiUtils.CreatePrettyError(c, 400, "какая-то ебатория с телеграммом", errors.New("err"))
			}
			newUser, err := userDtoFromUser(c, *userModel)
			if err != nil {
				return apiUtils.CreatePrettyError(c, 400, "какая-то ебатория с телеграммом", errors.New("err"))
			}
			return apiUtils.CreatePrettySuccess(c, newUser)
		}
	} else {
		if password == "" {
			// Анонимная регистрация
			newUser, err := anonymousRegister(c, username)
			if err != nil {
				return apiUtils.CreatePrettyError(c, 400, "какая-то ебатория с телеграммом", errors.New("err"))
			}
			return apiUtils.CreatePrettySuccess(c, newUser)
		} else {
			// проверяем парольную фразу
			newUser, err := anonymousLogin(c, username, password)
			if err != nil {
				return apiUtils.CreatePrettyError(c, 400, "какая-то ебатория с телеграммом", errors.New("err"))
			}
			return apiUtils.CreatePrettySuccess(c, newUser)
		}
	}
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
	tokenString := getIdFromToken(token)
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
