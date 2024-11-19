package web

type SendMessageDto struct {
	TelegramId int64  `json:"id"`
	Message    string `json:"message"`
}
