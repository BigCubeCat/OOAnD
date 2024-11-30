package groups

import (
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteGroup(c *fiber.Ctx) error {
	var group db.Group
	id, err := apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid id: "+strconv.Itoa(id), err)
	}
	dbInst := db.GetInstance()
	err = dbInst.First(&group, id).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "not found", err)
	}
	userId := userApi.GetCurrentUserId(c)
	if uint(userId) != group.OwnerID {
		return apiUtils.CreatePrettyError(c, 403, "forbidden", errors.New("forbidden"))
	}
	err = dbInst.Delete(&group, group.ID).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "invalid delete", err)
	}
	return apiUtils.CreatePrettySuccess(c, "")
}
