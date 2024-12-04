package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, START_MESSAGE)
	GetBotInstance().Send(msg)
}

func helpCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, HELP_MESSAGE)
	GetBotInstance().Send(msg)
}
