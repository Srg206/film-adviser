package senderbot

import (
	"film-adviser/repository"
	"film-adviser/settings"
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type SenderBot struct {
	token   string
	bot     *telego.Bot
	repo    *repository.Repository
	storage map[int64]string
}

func New() *SenderBot {
	return &SenderBot{}
}

func (sb *SenderBot) MustInit() {
	sb.token = settings.GetSettings().TgSenderToken
	var err error
	sb.bot, err = telego.NewBot(sb.token)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not start sender bot!")
	}
	//
	sb.storage = make(map[int64]string)
}

func (sb SenderBot) Handle() error {
	updates, _ := sb.bot.UpdatesViaLongPolling(nil)
	defer sb.bot.StopLongPolling()

	for update := range updates {
		fmt.Println(update)
		var chatID int64 // ID чата

		if update.Message != nil {
			chatID = update.Message.Chat.ID
			sb.storage[chatID] = update.Message.Text
			fmt.Println(sb.storage)
			// Сообщение
			message := tu.Message(
				tu.ID(chatID), // Используем правильный ID чата
				"Фильм успешно сохранен",
			)

			// Отправка сообщения
			_, _ = sb.bot.SendMessage(message)
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.GetChat().ID
		} else {
			continue
		}

	}
	return nil

}
