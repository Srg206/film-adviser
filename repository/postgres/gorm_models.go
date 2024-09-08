package postgres

import "gorm.io/gorm"

type UsersFilms struct {
	gorm.Model
	UserId int64
	Film   string
}

// table to link telegram user id with chat`s id
type UsersId struct {
	gorm.Model
	UserId     int64
	SenderId   int64
	ReceiverId int64
}
