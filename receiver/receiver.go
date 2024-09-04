package receiver

import "film-adviser/repository"

type Receiver interface {
	MustInit(repo repository.Repository)
	PickFilm(chatid int64) string
	SendAnswer()
}
