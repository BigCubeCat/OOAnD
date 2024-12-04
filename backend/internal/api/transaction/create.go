package transaction

import (
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"

	apiUtils "backend/internal/api/utils"
)

func CreateTransaction(c *fiber.Ctx) error {
	var (
		err         error
		transaction db.ClientTransactionRequest
	)
	err = c.BodyParser(&transaction)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid request Body", err)
	}
	err = db.GetInstance().Create(&transaction).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot create transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, transaction.ID)
}
