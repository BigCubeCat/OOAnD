package paymentmethod

import (
	"backend/internal/db"

	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetMyPaymentMethods(c *fiber.Ctx) error {
	var result []db.PaymentMethod
	userId := userApi.GetCurrentUserId(c)
	dbInst := db.GetInstance()
	if err := dbInst.Where("user_id = ?", userId).Find(&result).Error; err != nil {
		apiUtils.CreatePrettyError(c, 404, "not found", err)
	}
	return apiUtils.CreatePrettySuccess(c, result)
}

func GetPaymentMethods(c *fiber.Ctx) error {
	var result []db.PaymentMethod
	paramId, err := apiUtils.ParseId(c)
	if err != nil {
		apiUtils.CreatePrettyError(c, 400, "invalid id", err)
	}
	dbInst := db.GetInstance()
	if err := dbInst.Where("user_id = ?", paramId).Find(&result).Error; err != nil {
		apiUtils.CreatePrettyError(c, 404, "not found", err)
	}
	return apiUtils.CreatePrettySuccess(c, result)
}
