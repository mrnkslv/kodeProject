package handler

import (
	"encoding/json"
	"net/http"

	note "github.com/mrnkslv/kodeProject"
	"github.com/mrnkslv/kodeProject/pkg/service"
)

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	var input note.Note
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// yandex speller logic
	input.Description, err = service.SpellCheck(input.Description)
	if err != nil {
		panic(err)
	}

	input.Text, err = service.SpellCheck(input.Text)
	if err != nil {
		panic(err)
	}

	newNote, err := h.services.Note.Create(userId, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"new note": newNote,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type getAllNotesResponse struct {
	Data []note.Note `json:"data"`
}

func (h *Handler) getNotes(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	notes, err := h.services.Note.GetAll(userId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := getAllNotesResponse{
		Data: notes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
