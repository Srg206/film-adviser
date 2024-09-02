package receiver

type Receiver interface {
	MustInit()
	PickFilm(chatid int64) string
	SendAnswer()
}
