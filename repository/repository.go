package repository

type Repository interface {
	MustInit()
	Write(chatid int64, film string) error
	PickRandom(chatid int64) (error, string)
}
