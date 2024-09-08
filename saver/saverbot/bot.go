package saverbot

import (
	"film-adviser/repository"
	"film-adviser/settings"
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type SaverBot struct {
	token string
	bot   *telego.Bot
	repo  repository.Repository
}

func New() *SaverBot {
	return &SaverBot{}
}

// func to initialise saver bot
func (sb *SaverBot) MustInit(repo repository.Repository) {
	sb.token = settings.GetSettings().TgSaverToken
	var err error
	sb.bot, err = telego.NewBot(sb.token)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not start sender bot!")
	}
	sb.repo = repo
}

// func to start saver bot
func (sb SaverBot) Handle() error {
	updates, _ := sb.bot.UpdatesViaLongPolling(nil)
	defer sb.bot.StopLongPolling()

	for update := range updates {
		if update.Message.Text == "/start" {
			sb.repo.AddChatid(0, update.Message.Chat.ID, update.Message.From.ID)
			continue
		}
		//sb.repo.AddChatid(0, update.Message.Chat.ID, update.Message.From.ID)

		fmt.Println(update)
		var chatID int64

		if update.Message != nil {
			sb.repo.Write(update.Message.From.ID, update.Message.Text)

			chatID = update.Message.Chat.ID
			message := tu.Message(
				tu.ID(chatID),
				"Film saved successfully",
			)

			_, _ = sb.bot.SendMessage(message)
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.GetChat().ID
		} else {
			continue
		}

	}
	return nil
}
