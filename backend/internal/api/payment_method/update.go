package paymentmethod

import (
	"backend/internal/db"

	"backend/internal/api/dto"
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"

	"github.com/gofiber/fiber/v2"
)

func PatchPaymentMethod(c *fiber.Ctx) error {
	var (
		paramId int
		err     error
		dto     dto.PaymentMethodDto
		method  db.PaymentMethod
	)
	paramId, err = apiUtils.ParseId(c)
	if err != nil {

	}
	dbInst := db.GetInstance()
	if err := dbInst.Where("payment_id = ?", paramId).First(&method).Error; err != nil {
		return apiUtils.CreatePrettyError(c, 404, "not found", err)
	}
	if method.UserID != userApi.GetCurrentUserId(c) {
		return apiUtils.CreatePrettyError(c, 403, "forbidden", err)
	}
	err = c.BodyParser(&dto)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid request Body", err)
	}
	method.UserID = userApi.GetCurrentUserId(c)
	if dto.Name != "" {
		method.Name = dto.Name
	}
	if dto.Requisites != "" {
		method.Requisites = dto.Requisites
	}
	if dto.Type != "" {
		method.Type = dto.Type
	}
	if dbInst.Save(&method).Error != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot delete method", err)
	}
	return apiUtils.CreatePrettySuccess(c, method.PaymentId)
}
