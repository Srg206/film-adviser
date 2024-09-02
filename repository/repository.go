package repository

type Repository interface {
	Write() error
	Read() error
}
