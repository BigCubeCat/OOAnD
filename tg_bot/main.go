package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // Для отладки
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Запускаем HTTP-сервер
	http.HandleFunc("/user/", handleUserMessage)
	go func() {
		log.Println("Запуск HTTP-сервера на порту 8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Обработка команд бота
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
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Неизвестная команда. Используй /help для списка доступных команд.",
			)
			bot.Send(msg)
		}
	}
}

func startCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		"Привет! Я бот, который может отправлять сообщения. "+
			"Используй HTTP POST запросы для отправки сообщений или команды в Telegram.",
	)
	bot.Send(msg)
}

func helpCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Доступные команды:\n"+
		"/start - начать работу с ботом\n"+
		"/help - показать это сообщение")
	bot.Send(msg)
}

// Обработчик для POST запросов на /user/{telegram id}/{message}
func handleUserMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы поддерживаются", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем параметры из URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(
			w,
			"Неверный URL. Используйте формат /user/{telegram_id}/{message}",
			http.StatusBadRequest,
		)
		return
	}

	// Извлекаем Telegram ID
	telegramID, err := strconv.ParseInt(pathParts[2], 10, 64)
	if err != nil {
		http.Error(w, "Неверный Telegram ID", http.StatusBadRequest)
		return
	}

	// Извлекаем сообщение
	message := strings.Join(pathParts[3:], "/")
	if message == "" {
		http.Error(w, "Сообщение не может быть пустым", http.StatusBadRequest)
		return
	}

	// Отправляем сообщение пользователю
	msg := tgbotapi.NewMessage(telegramID, message)
	if _, err := bot.Send(msg); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Ошибка отправки сообщения: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Ответ на успешный запрос
	response := map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Сообщение успешно отправлено пользователю %d", telegramID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
