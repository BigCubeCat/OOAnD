package groups

import (
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
)

func GetMyGroups(c *fiber.Ctx) error {
	var groups []db.Group
	userID := userApi.GetCurrentUserId(c)
	dbInst := db.GetInstance()
	// Запрос на получение пользователя и связанных групп.
	err := dbInst.Where("owner_id = ?", userID).Find(&groups).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "some error", err)
	}
	return apiUtils.CreatePrettySuccess(c, groups)
}
