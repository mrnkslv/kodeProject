package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	note "github.com/mrnkslv/kodeProject"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (r *NotePostgres) Create(userId int, note note.Note) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createNoteQuery := fmt.Sprintf("INSERT INTO %s (text,description) VALUES ($1,$2) RETURNING id", notesTable)
	row := tx.QueryRow(createNoteQuery, note.Text, note.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserNoteQuery := fmt.Sprintf("INSERT INTO %s (user_id,note_id) VALUES ($1,$2)", usersNotesTable)
	_, err = tx.Exec(createUserNoteQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()

}

func (r *NotePostgres) GetAll(userId int) ([]note.Note, error) {
	var notes []note.Note
	query := fmt.Sprintf("SELECT n.id, n.text, n.description  FROM %s n INNER JOIN %s un on n.id=un.note_id WHERE un.user_id = $1", notesTable, usersNotesTable)
	err := r.db.Select(&notes, query, userId)

	return notes, err
}
