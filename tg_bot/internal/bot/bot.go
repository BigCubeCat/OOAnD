package bot

import (
	"log"
	"sync"
	"tg_bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	once sync.Once
	bot  *tgbotapi.BotAPI
)

func GetBotInstance() *tgbotapi.BotAPI {
	once.Do(func() {
		var err error
		bot = new(tgbotapi.BotAPI)
		bot, err = tgbotapi.NewBotAPI(config.GetConfig().TgBotToken)
		if err != nil {
			log.Println("cant create BOT")
			log.Panic(err)
		}

		bot.Debug = true // Для отладки
		log.Printf("Authorized on account %s", bot.Self.UserName)
	})
	return bot
}

func ServeBot() {
	bot := GetBotInstance()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "start":
			startCommand(update.Message)
		case "help":
			helpCommand(update.Message)
		default:
			if err := SendMessage(update.Message.Chat.ID, HELP_MESSAGE); err != nil {
				log.Println("error: ", err.Error())
			}
		}
	}
}

func SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	return err
}
