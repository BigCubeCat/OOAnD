package dto

type UserAccountDto struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	TelegramId int    `json:"telegram_id"`
	Token      string `json:"token"`
}
