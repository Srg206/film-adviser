package reminderbot

import tu "github.com/mymmrac/telego/telegoutil"

var (
	inlineKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Remind me film").
				WithCallbackData("remind_film"),
		),
	)
)
