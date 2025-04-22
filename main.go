package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не установлен")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Обработка сообщений
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Напиши /menu для демонстрации inline-кнопок.")
				bot.Send(msg)

			case "/menu":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери вариант:")

				buttons := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("🔍 Узнать больше", "more_info"),
						tgbotapi.NewInlineKeyboardButtonData("❌ Закрыть", "delete_msg"),
					),
				)

				msg.ReplyMarkup = buttons
				bot.Send(msg)
			}
		}

		// Обработка нажатий на inline-кнопки
		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data

			switch data {
			case "more_info":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Интересный факт: Telegram API — очень мощный инструмент! 🚀")
				bot.Send(msg)

			case "delete_msg":
				del := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				bot.Send(del)
			}

			// Уведомим Telegram, что мы обработали callback
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			bot.Request(callback)
		}
	}
}
