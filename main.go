package main

import (
	"film-adviser/settings"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"golang.org/x/exp/rand"
)

var storage []string

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	fmt.Println("From singleton token      " + settings.GetSettings().TgSenderToken)
	fmt.Println("From singleton token      " + settings.GetSettings().TgSenderToken)
	fmt.Println("From singleton token      " + settings.GetSettings().TgSenderToken)
	fmt.Println("From singleton token      " + settings.GetSettings().TgSenderToken)
	fmt.Println("From singleton token      " + settings.GetSettings().TgSenderToken)

	botToken := os.Getenv("TG_BOT_TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Inline keyboard parameters
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 1
			tu.InlineKeyboardButton("Сохранить фильм").
				WithCallbackData("save_film"),
			tu.InlineKeyboardButton("Порекомендуй фильм").
				WithCallbackData("recomend_film"),
		),
	)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		var chatID int64 // ID чата

		if update.Message != nil {
			chatID = update.Message.Chat.ID
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.GetChat().ID
		} else {
			continue
		}

		if update.CallbackQuery == nil {
			filmName := update.Message.Text
			fmt.Println("Пользователь ввел название фильма:", filmName)
			storage = append(storage, filmName)
			saved_film_msg := tu.Message(
				tu.ID(chatID),
				"Фильм успешно сохранен",
			)
			_, _ = bot.SendMessage(saved_film_msg)
		}

		// Определяем, на какую кнопку нажал пользователь
		if update.CallbackQuery != nil {
			callbackData := update.CallbackQuery.Data

			// Определяем, на какую кнопку нажал пользователь
			switch callbackData {
			case "save_film":
				// Отправляем сообщение с просьбой ввести название фильма
				say_film_msg := tu.Message(
					tu.ID(chatID), // Используем правильный ID чата
					"Введите название фильма",
				)
				_, _ = bot.SendMessage(say_film_msg)

				// Ждем, пока пользователь введет название фильма
				for update := range updates {
					if update.Message != nil && update.Message.Chat.ID == chatID {
						filmName := update.Message.Text
						fmt.Println("Пользователь ввел название фильма:", filmName)
						storage = append(storage, filmName)
						saved_film_msg := tu.Message(
							tu.ID(chatID),
							"Фильм успешно сохранен",
						)
						_, _ = bot.SendMessage(saved_film_msg)
						break // Выходим из внутреннего цикла
					}
				}

			case "recomend_film":
				fmt.Println("Пользователь нажал на кнопку 'Порекомендуй фильм'")
				rand_id := rand.Intn(len(storage))
				film_msg := tu.Message(
					tu.ID(chatID),
					storage[rand_id],
				)
				_, _ = bot.SendMessage(film_msg)
			default:
				fmt.Println("Неизвестная кнопка")
			}
		}

		// Сообщение
		message := tu.Message(
			tu.ID(chatID), // Используем правильный ID чата
			"My message",
		).WithReplyMarkup(inlineKeyboard)

		// Отправка сообщения
		_, _ = bot.SendMessage(message)
	}

	// Receiving callback data
}
