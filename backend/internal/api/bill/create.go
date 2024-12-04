package bill

import (
	"backend/internal/api/dto"
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
)

func CreateBill(c *fiber.Ctx) error {
	dto := new(dto.CreateBillDTO)
	if err := c.BodyParser(dto); err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Invalid request Body", err)
	}
	bill := new(db.Bill)
	bill.Name = dto.Name
	owner, err := userApi.GetCurrentUser(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "cant get current user", err)
	}
	bill.Owner = owner.SerialID
	for _, pos := range dto.Positions {
		bill.BillPositions = append(bill.BillPositions, db.BillPosition{
			Name:         pos.Name,
			WhoPaid:      pos.WhoPaid,
			FromWhomPaid: pos.FromWhomPaid,
			Amount:       pos.Amount,
		})
	}
	err = db.GetInstance().Create(bill).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Invalid Create", err)
	}
	return apiUtils.CreatePrettySuccess(c, bill.ID)
}
