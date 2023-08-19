package service

import (
	note "github.com/mrnkslv/kodeProject/models"
	"github.com/mrnkslv/kodeProject/pkg/repository"
)

type Authorization interface {
	CreateUser(user note.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Note interface {
	Create(userId int, note note.Note) (note.Note, error)
	GetAll(userId int) ([]note.Note, error)
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Note:          NewNotesService(repos.Note),
	}
}
