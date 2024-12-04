package web

import (
	"log"
	"tg_bot/internal/bot"

	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		var (
			message SendMessageDto
			err     error
		)
		err = c.BodyParser(&message)
		if err != nil {
			return c.Status(400).
				JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
		}
		err = bot.SendMessage(message.TelegramId, message.Message)
		if err != nil {
			return c.Status(400).
				JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
		}
		return c.JSON(fiber.Map{"status": "success", "data": nil})
	})

	log.Fatal(app.Listen(":7777"))
}
