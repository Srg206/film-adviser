package repository

type Repository interface {
	MustInit()
	Write(chatid int64, film string) error
	PickRandom(chatid int64) (error, string)
	AddChatid(receiver_id, sender_id, user_id int64)
}
