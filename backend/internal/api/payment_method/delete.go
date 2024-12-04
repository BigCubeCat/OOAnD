package paymentmethod

import (
	"backend/internal/db"

	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"

	"github.com/gofiber/fiber/v2"
)

func DeletePaymentMethod(c *fiber.Ctx) error {
	var (
		paramId int
		err     error
		method  db.PaymentMethod
	)
	paramId, err = apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid method id", err)
	}

	dbInst := db.GetInstance()
	err = dbInst.Where("id = ?", paramId).First(&method).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "not found", err)
	}
	if method.UserID != userApi.GetCurrentUserId(c) {
		return apiUtils.CreatePrettyError(c, 403, "forbidden", err)
	}
	err = dbInst.Delete(&method).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot delete method", err)
	}
	return apiUtils.CreatePrettySuccess(c, method.PaymentId)
}
