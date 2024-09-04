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
	token string
	bot   *telego.Bot
	repo  repository.Repository
}

func New() *SenderBot {
	return &SenderBot{}
}

func (sb *SenderBot) MustInit(repo repository.Repository) {
	sb.token = settings.GetSettings().TgSenderToken
	var err error
	sb.bot, err = telego.NewBot(sb.token)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not start sender bot!")
	}
	sb.repo = repo
	//
}

func (sb SenderBot) Handle() error {
	updates, _ := sb.bot.UpdatesViaLongPolling(nil)
	defer sb.bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start" {
			//update.Message.From.ID
		}
		fmt.Println(update)
		var chatID int64 // ID чата

		if update.Message != nil {
			sb.repo.Write(update.Message.Chat.ID, update.Message.Text)

			chatID = update.Message.Chat.ID
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

//func (sb SenderBot) add If
