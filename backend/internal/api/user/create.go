package user

import (
	"backend/internal/api/dto"
	"backend/internal/db"
	"backend/internal/utils"
	"errors"

	apiUtils "backend/internal/api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
// @Summary Создание нового пользователя
// @Description Создает нового пользователя либо по telegram_id либо по email
// @Router /api/user [post]
func CreateUser(
	c *fiber.Ctx,
	telegram_id int,
	lastname string,
	firstname string,
	username string,
	avatar string,
) error {
	user := new(db.User)

	user.TelegramID = telegram_id
	user.LastName = lastname
	user.FirstName = firstname
	user.Username = username
	user.Avatar = avatar

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
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		TelegramId: user.TelegramID,
		Token:      t,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

func userDtoFromUser(c *fiber.Ctx, user db.User) (dto.UserAccountDto, error) {
	var newUser dto.UserAccountDto
	t, err := createToken(user.SerialID)
	if err != nil {
		return newUser, apiUtils.CreatePrettyError(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
			err,
		)
	}

	return dto.UserAccountDto{
		Id:         user.SerialID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		TelegramId: user.TelegramID,
		Token:      t,
	}, nil
}

func anonymousRegister(c *fiber.Ctx, username string) (dto.UserAccountDto, error) {
	var newUser db.User
	var newUserDto dto.UserAccountDto
	dbInst := db.GetInstance()
	err := dbInst.First(&newUser).Where("username = ?", username).Error
	if err != nil {
		return newUserDto, err
	}
	if newUser.SerialID > 0 {
		errorMessage := "user " + username + " already exists"
		return newUserDto, errors.New(errorMessage)
	}
	newUser.Password = utils.GenerateRandomPassword()
	err = dbInst.Create(&newUser).Error
	if err != nil {
		return newUserDto, err
	}
	return userDtoFromUser(c, newUser)
}

func anonymousLogin(c *fiber.Ctx, username string, password string) (dto.UserAccountDto, error) {
	var newUser db.User
	var newUserDto dto.UserAccountDto
	dbInst := db.GetInstance()
	err := dbInst.First(&newUser).Where("username = ?", username).Error
	if err != nil {
		return newUserDto, err
	}
	if newUser.SerialID == 0 {
		errorMessage := "user " + username + " not exists"
		return newUserDto, errors.New(errorMessage)
	}
	if newUser.Password == password {
		return userDtoFromUser(c, newUser)
	}
	return newUserDto, errors.New("bad path phrase")
}
