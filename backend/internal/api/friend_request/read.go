package friendrequest

import (
	"backend/internal/config"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"
)

func GetMyFriendRequests(c *fiber.Ctx) error {
	var requests []db.FriendRequest
	userId := userApi.GetCurrentUserId(c)
	logrus.Error(userId)
	err := db.GetInstance().
		Preload("FromUser").
		Where(
			"to_user_id = ? AND state = ? OR state = ?",
			userId,
			config.PENDING_FRIEND_STATE,
			config.REJECTED_FRIEND_STATE,
		).Find(&requests).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 404, "not found", err)
	}
	return apiUtils.CreatePrettySuccess(c, requests)
}
