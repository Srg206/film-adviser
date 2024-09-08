package reminderbot

import (
	"film-adviser/repository"
	"film-adviser/settings"
	"fmt"
	"log"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type RecomendBot struct {
	bot   *telego.Bot
	token string
	repo  repository.Repository
}

func New() *RecomendBot {
	return &RecomendBot{}
}

func (rb *RecomendBot) MustInit(repo repository.Repository) {
	rb.token = settings.GetSettings().TgReminderToken
	var err error
	rb.bot, err = telego.NewBot(rb.token)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not start sender bot!")
	}
	rb.repo = repo
}

func (rb RecomendBot) PickFilm(chatid int64) string {

	if err, res := rb.repo.PickRandom(chatid); err == nil {
		return res
	} else {
		fmt.Println("Could not pick film !")
		return ""
	}
}

func (rb RecomendBot) SendAnswer() {
	StartTime := time.Now().Unix()
	updates, _ := rb.bot.UpdatesViaLongPolling(&telego.GetUpdatesParams{
		Timeout: 1,
	})
	defer rb.bot.StopLongPolling()

	for update := range updates {
		// check time in order to do not process old updates
		var updateTime int64
		if update.Message != nil {
			updateTime = update.Message.GetDate()
		}
		if update.CallbackQuery != nil {
			updateTime = update.CallbackQuery.Message.GetDate()
		}

		if updateTime < StartTime {
			continue
		}

		// process start
		if update.Message != nil && update.Message.Text == "/start" {
			rb.repo.AddChatid(update.Message.Chat.ID, 0, update.Message.From.ID)
			message := tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Let me remind you of the film",
			).WithReplyMarkup(inlineKeyboard)

			_, _ = rb.bot.SendMessage(message)
			continue
		}

		// define user id
		var UserID int64
		var chatID int64
		if update.Message != nil {
			UserID = update.Message.From.ID
			_, chatID, _ = rb.repo.GetUserChat(UserID)
		} else if update.CallbackQuery != nil {
			UserID = update.CallbackQuery.From.ID
			_, chatID, _ = rb.repo.GetUserChat(UserID)

		} else {
			continue
		}

		// process clicking remind_film
		if update.CallbackQuery != nil {
			callbackData := update.CallbackQuery.Data
			if callbackData == "remind_film" {
				message := tu.Message(
					tu.ID(chatID),
					rb.PickFilm(chatID),
				).WithReplyMarkup(inlineKeyboard)

				_, _ = rb.bot.SendMessage(message)
			}
		} else {
			message := tu.Message(
				tu.ID(chatID),
				"Let me remind you of the film",
			).WithReplyMarkup(inlineKeyboard)

			_, _ = rb.bot.SendMessage(message)
		}

	}

}
