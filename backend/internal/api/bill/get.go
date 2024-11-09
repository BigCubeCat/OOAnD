package bill

import (
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBill(c *fiber.Ctx) error {
	var (
		idParam string
		id      int
		err     error
		bill    db.Bill
	)
	// TODO: проверка что пользователь именно тот, что указан в чеке
	idParam = c.Params("id")
	id, err = strconv.Atoi(idParam)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid Id: "+idParam, err)
	}
	err = db.GetInstance().First(&bill, id).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "bill not found", err)
	}
	return apiUtils.CreatePrettySuccess(c, bill)
}

func GetAllUserBills(c *fiber.Ctx) error {
	// TODO: Выбрать только МОИ чеки
	var (
		err      error
		page     int
		pageSize int
		allBills []db.Bill
	)
	query := c.Queries()
	pageParam := query["page"]
	pageSizeParam := query["size"]
	if pageSizeParam == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(pageSizeParam)
		if err != nil {
			return apiUtils.CreatePrettyError(c, 400, "Invalid page size: "+pageSizeParam, err)
		}
	}
	page, err = strconv.Atoi(pageParam)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid page number: "+pageParam, err)
	}

	offset := (page - 1) * pageSize
	err = db.GetInstance().Limit(pageSize).Offset(offset).Find(&allBills).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "page "+pageParam+" not found", err)
	}
	return apiUtils.CreatePrettySuccess(c, allBills)
}
