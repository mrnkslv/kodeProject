package handler

import (
	"encoding/json"
	"net/http"

	note "github.com/mrnkslv/kodeProject/models"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input note.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"id": id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
