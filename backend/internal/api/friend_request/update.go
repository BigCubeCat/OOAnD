package friendrequest

import (
	apiUtils "backend/internal/api/utils"
	"backend/internal/config"
	"backend/internal/db"

	"github.com/gofiber/fiber/v2"
)

func UpdateFriendRequest(c *fiber.Ctx) error {
	var (
		err           error
		friendRequest db.FriendRequest
	)
	err = c.BodyParser(&friendRequest)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid request Body", err)
	}
	id, err := apiUtils.ParseId(c)
	friendRequest.ID = uint(id)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "Invalid Id", err)
	}
	err = db.GetInstance().Save(&friendRequest).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot create transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, friendRequest.ID)
}

func AcceptFriendRequest(c *fiber.Ctx) error {
	var (
		err error
		req db.FriendRequest
	)
	id, err := apiUtils.ParseId(c)
	req.ID = uint(id)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "invalid id", err)
	}
	err = db.GetInstance().
		Model(&req).
		Select("State").
		Updates(db.ClientTransactionRequest{State: config.ACCEPT_FRIEND_STATE}).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot accept transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, req)
}

func DeclineFriendRequest(c *fiber.Ctx) error {
	var (
		err error
		req db.FriendRequest
	)
	id, err := apiUtils.ParseId(c)
	req.ID = uint(id)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 400, "invalid id", err)
	}
	err = db.GetInstance().
		Model(&req).
		Select("State").
		Updates(db.ClientTransactionRequest{State: config.REJECTED_FRIEND_STATE}).Error
	if err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Cannot accept transaction", err)
	}
	return apiUtils.CreatePrettySuccess(c, req)
}
