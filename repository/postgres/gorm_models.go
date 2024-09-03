package postgres

import "gorm.io/gorm"

type UsersFilms struct {
	gorm.Model
	Chatid int64
	Film   string
}
