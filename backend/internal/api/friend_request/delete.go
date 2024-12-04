package friendrequest

import (
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"

	apiUtils "backend/internal/api/utils"
)

func DeleteFriendRequest(c *fiber.Ctx) error {
	var (
		err           error
		friendRequest db.FriendRequest
	)
	id, err := apiUtils.ParseId(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "invalid id", err)
	}
	err = db.GetInstance().Delete(&db.FriendRequest{ID: uint(id)}).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "cannot delete transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, friendRequest.ID)
}
