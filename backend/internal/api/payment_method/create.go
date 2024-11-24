package paymentmethod

import (
	"backend/internal/db"

	"backend/internal/api/dto"
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"

	"github.com/gofiber/fiber/v2"
)

func AddPaymentMethod(c *fiber.Ctx) error {
	var (
		err       error
		methodDto dto.PaymentMethodDto
	)
	err = c.BodyParser(&methodDto)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid request Body", err)
	}
	var method db.PaymentMethod
	method.UserID = userApi.GetCurrentUserId(c)
	method.Name = methodDto.Name
	method.Requisites = methodDto.Requisites
	method.Type = methodDto.Type

	err = db.GetInstance().Create(&method).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot create transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, method.PaymentId)
}
