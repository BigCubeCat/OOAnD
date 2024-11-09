package bill

import (
	"backend/internal/api/dto"
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
)

func UpdateBill(c *fiber.Ctx) error {
	var (
		idParam string
		id      int
		err     error
		bill    db.Bill
	)
	dbInst := db.GetInstance()
	// TODO: проверка что пользователь именно тот, что указан в чеке
	id, err = apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid Id: "+idParam, err)
	}
	dto := new(dto.CreateBillDTO)
	if err := c.BodyParser(dto); err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Invalid request Body", err)
	}
	err = dbInst.Preload("BillPositions").First(&bill, id).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "bill not found", err)
	}
	bill.Name = dto.Name
	bill.Owner = 1 // TODO: get Owner from context
	for _, pos := range dto.Positions {
		bill.BillPositions = append(bill.BillPositions, db.BillPosition{
			Name:         pos.Name,
			WhoPaid:      pos.WhoPaid,
			FromWhomPaid: pos.FromWhomPaid,
			Amount:       pos.Amount,
		})
	}

	var updatedPositions []db.BillPosition
	for _, newPos := range bill.BillPositions {
		var position db.BillPosition

		// Если позиция с ID существует, обновляем ее
		if newPos.ID != 0 {
			if err := dbInst.First(&position, newPos.ID).Error; err == nil {
				position.Name = newPos.Name
				position.WhoPaid = newPos.WhoPaid
				position.FromWhomPaid = newPos.FromWhomPaid
				position.Amount = newPos.Amount
				dbInst.Save(&position)
			}
		} else {
			// Добавляем новую позицию
			newPos.IDBill = bill.ID
			updatedPositions = append(updatedPositions, newPos)
		}
	}
	bill.BillPositions = updatedPositions

	err = dbInst.Model(&bill).Association("BillPositions").Replace(updatedPositions)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "bill not found", err)
	}
	err = dbInst.Save(&bill).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "error on update", err)
	}
	return apiUtils.CreatePrettySuccess(c, bill)
}
