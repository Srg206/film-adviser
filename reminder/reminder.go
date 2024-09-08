package reminder

import "film-adviser/repository"

type Reminder interface {
	MustInit(repo repository.Repository)
	PickFilm(chatid int64) string
	SendAnswer()
}
