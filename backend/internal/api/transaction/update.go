package transaction

import (
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
)

func UpdateTransaction(c *fiber.Ctx) error {
	var (
		err         error
		transaction db.ClientTransactionRequest
	)
	err = c.BodyParser(&transaction)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid request Body", err)
	}
	transaction.ID, err = apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid Id", err)
	}
	err = db.GetInstance().Save(&transaction).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot create transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, transaction.ID)
}

func AcceptTransaction(c *fiber.Ctx) error {
	var (
		err         error
		transaction db.ClientTransactionRequest
	)
	transaction.ID, err = apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "invalid id", err)
	}
	err = db.GetInstance().
		Model(&transaction).
		Select("State").
		Updates(db.ClientTransactionRequest{State: "accepted"}).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot accept transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, transaction)
}

func DeclineTransaction(c *fiber.Ctx) error {
	var (
		err         error
		transaction db.ClientTransactionRequest
	)
	transaction.ID, err = apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "invalid id", err)
	}
	err = db.GetInstance().
		Model(&transaction).
		Select("State").
		Updates(db.ClientTransactionRequest{State: "declined"}).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot accept transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, transaction)
}
