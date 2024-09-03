package receiver

import "film-adviser/repository"

type Receiver interface {
	MustInit()
	PickFilm(repo *repository.Repository, chatid int64) string
	SendAnswer()
}
