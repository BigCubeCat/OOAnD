package main

import (
	"tg_bot/internal/bot"
	"tg_bot/internal/web"
)

func main() {
	go bot.ServeBot()
	web.Serve()
}
