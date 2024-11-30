package groups

import (
	"backend/internal/api/dto"
	userApi "backend/internal/api/user"
	apiUtils "backend/internal/api/utils"
	"backend/internal/db"
	_ "strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateGroup(c *fiber.Ctx) error {
	dto := new(dto.CreateGroupDTO)
	if err := c.BodyParser(dto); err != nil {
		return apiUtils.CreatePrettyError(c, 500, "Invalid request Body", err)
	}
	group := new(db.Group)
	currentUser, err := userApi.GetCurrentUser(c)
	if err != nil {
		return apiUtils.CreatePrettyError(c, 403, "Invalid user", err)
	}
	group.Owner = *currentUser
	group.Name = dto.Name
	group.Description = dto.Description
	dbInst := db.GetInstance()
	dbInst.Create(&group)
	//
	// err = dbInst.Create(&db.UsersGroups{User: *currentUser, Group: *group}).Error
	// if err != nil {
	// 	return apiUtils.CreatePrettyError(c, 500, "Invalid create", err)
	// }
	// var user db.User
	// for _, member := range dto.Members {
	// 	err = dbInst.First(&user, member).Error
	// 	if err != nil {
	// 		return apiUtils.CreatePrettyError(c, 400, "Invalid user id: "+strconv.Itoa(member), err)
	// 	}
	// 	err = dbInst.Create(&db.UsersGroups{User: user, Group: *group}).Error
	// 	if err != nil {
	// 		return apiUtils.CreatePrettyError(c, 500, "Invalid create relation", err)
	// 	}
	// }
	// TODO: проверка на друзей
	return apiUtils.CreatePrettySuccess(c, group.ID)
}
