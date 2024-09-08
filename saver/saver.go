package saver

import "film-adviser/repository"

// Saver is component to write saved films into repository

type Saver interface {
	MustInit(repo repository.Repository)
	Handle() error
}
