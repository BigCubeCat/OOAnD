package transaction

import (
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"

	apiUtils "backend/internal/api/utils"
)

func DeleteTransaction(c *fiber.Ctx) error {
	var (
		err         error
		transaction db.ClientTransactionRequest
	)
	id, err := apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "invalid id", err)
	}
	err = db.GetInstance().Delete(&db.ClientTransactionRequest{ID: id}).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "cannot delete transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, transaction.ID)
}
