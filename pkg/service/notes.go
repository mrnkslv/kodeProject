package service

import (
	note "github.com/mrnkslv/kodeProject/models"
	"github.com/mrnkslv/kodeProject/pkg/repository"
)

type NotesService struct {
	repo repository.Note
}

func NewNotesService(repo repository.Note) *NotesService {
	return &NotesService{repo: repo}
}

func (s *NotesService) Create(userId int, note note.Note) (note.Note, error) {
	return s.repo.Create(userId, note)
}

func (s *NotesService) GetAll(userId int) ([]note.Note, error) {
	return s.repo.GetAll(userId)
}
