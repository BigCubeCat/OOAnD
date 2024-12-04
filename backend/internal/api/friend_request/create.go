package friendrequest

import (
	"backend/internal/config"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"

	"backend/internal/api/dto"
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"
)

func CreateFriendRequest(c *fiber.Ctx) error {
	var (
		err           error
		dto           dto.FriendRequestDTO
		friendRequest db.FriendRequest
	)
	err = c.BodyParser(&dto)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid request Body", err)
	}
	userPtr, err := userApi.GetCurrentUser(c)
	otherUser, err := userApi.GetUserById(dto.To)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "invalid user id", err)
	}

	friendRequest.FromUserID = uint(userPtr.SerialID)
	friendRequest.ToUserID = uint(otherUser.SerialID)
	friendRequest.State = config.PENDING_FRIEND_STATE

	err = db.GetInstance().Create(&friendRequest).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot create friends request", err)
	}
	return apiUtils.CreatePrettySuccess(c, friendRequest.ID)
}
