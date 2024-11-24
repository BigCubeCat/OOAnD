package bill

import (
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func DeleteBill(c *fiber.Ctx) error {
	var (
		bill db.Bill
	)
	id, err := apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid Id", err)
	}
	dbInst := db.GetInstance()
	err = dbInst.Preload("BillPositions").
		Where("id = ?", id).
		First(&bill).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "bill not found", err)
	}
	log.Debug(bill)
	currentUser, err := userApi.GetCurrentUser(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 403, "forbidden", err)
	}
	if bill.Owner != currentUser.SerialID {
		return apiUtils.CreatePrettyError(c, 403, "forbidden", err)
	}
	for _, pos := range bill.BillPositions {
		err = dbInst.Delete(&pos).Error
		if err != nil {
			return apiUtils.CreatePrettyError(c, 500, "cannot delete", err)
		}
	}
	err = dbInst.Delete(&bill).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "cannot delete", err)
	}
	return apiUtils.CreatePrettySuccess(c, "success")
}
