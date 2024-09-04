package repository

type Repository interface {
	MustInit()
	Write(userid int64, film string) error
	PickRandom(chatid int64) (error, string)
	AddChatid(receiver_id, sender_id, user_id int64) error
	GetUserChat(user_id int64) (error, int64, int64)
}
