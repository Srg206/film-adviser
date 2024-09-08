package repository

type Repository interface {
	MustInit()
	Write(userid int64, film string) error
	PickRandom(chatid int64) (error, string)

	// add id of chat in telegram to table to link id of saverbot chat and reminder chat
	AddChatid(reminder_id, saver_id, user_id int64) error
	GetUserChat(user_id int64) (error, int64, int64)
}
