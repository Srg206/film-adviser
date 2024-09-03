package sender

import "film-adviser/repository"

// Sender is component to write saved films into repository

type Sender interface {
	MustInit(repo repository.Repository)
	Handle() error
}
