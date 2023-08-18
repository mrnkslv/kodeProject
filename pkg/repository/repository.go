package repository

import (
	"github.com/jmoiron/sqlx"
	note "github.com/mrnkslv/kodeProject"
)

type Authorization interface {
	CreateUser(user note.User) (int, error)
	GetUser(username, password string) (note.User, error)
}

type Note interface {
	Create(userId int, note note.Note) (int, error)
	GetAll(userId int) ([]note.Note, error)
}

type Repository struct {
	Authorization
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Note:          NewNotePostgres(db),
	}
}
