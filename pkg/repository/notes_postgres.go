package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	note "github.com/mrnkslv/kodeProject/models"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (r *NotePostgres) Create(userId int, newNote note.Note) (note.Note, error) {
	createNoteQuery := fmt.Sprintf("INSERT INTO %s (text,description,user_id) VALUES ($1,$2,$3) RETURNING id,user_id,text,description", notesTable)
	row := r.db.QueryRow(createNoteQuery, newNote.Text, newNote.Description, userId)
	if err := row.Scan(&newNote.Id, &newNote.UserId, &newNote.Text, &newNote.Description); err != nil {
		return newNote, err
	}

	return newNote, nil

}

func (r *NotePostgres) GetAll(userId int) ([]note.Note, error) {
	var notes []note.Note
	getQuery := fmt.Sprintf("SELECT id, text, description FROM %s WHERE user_id = $1", notesTable)
	err := r.db.Select(&notes, getQuery, userId)
	return notes, err
}
