package receiverbot

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
	rb.token = settings.GetSettings().TgReceiverToken
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
	fmt.Println("m\n\n\n\n\n\nemgkmrkg\n\nm")
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 1
			tu.InlineKeyboardButton("Порекомендуй фильм").
				WithCallbackData("recomend_film"),
		),
	)

	StartTime := time.Now().Unix()
	updates, _ := rb.bot.UpdatesViaLongPolling(nil)
	defer rb.bot.StopLongPolling()

	for update := range updates {
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

		if update.Message != nil && update.Message.Text == "/start" {
			rb.repo.AddChatid(update.Message.Chat.ID, 0, update.Message.From.ID)
			message := tu.Message(
				tu.ID(update.Message.Chat.ID), // Используем правильный ID чата
				"Давайте порекомендую вам фильм",
			).WithReplyMarkup(inlineKeyboard)

			// Отправка сообщения
			_, _ = rb.bot.SendMessage(message)
			continue
		}
		var UserID int64 // ID чата
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
		if update.CallbackQuery != nil {
			callbackData := update.CallbackQuery.Data
			if callbackData == "recomend_film" {
				message := tu.Message(
					tu.ID(chatID), // Используем правильный ID чата
					rb.PickFilm(chatID),
				).WithReplyMarkup(inlineKeyboard)

				// Отправка сообщения
				_, _ = rb.bot.SendMessage(message)
			}
		} else {
			message := tu.Message(
				tu.ID(chatID), // Используем правильный ID чата
				"Давайте порекомендую вам фильм",
			).WithReplyMarkup(inlineKeyboard)

			// Отправка сообщения
			_, _ = rb.bot.SendMessage(message)
		}

	}

}
