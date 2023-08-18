package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	note "github.com/mrnkslv/kodeProject"
	"github.com/mrnkslv/kodeProject/pkg/service"
)

func (h *Handler) createNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input note.Note
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	////yandex speller logic
	input.Description, err = service.SpellCheck(input.Description)
	if err != nil {
		panic(err)
	}

	input.Text, err = service.SpellCheck(input.Text)
	if err != nil {
		panic(err)
	}

	id, err := h.services.Note.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllNotesResponse struct {
	Data []note.Note `json:"data"`
}

func (h *Handler) getNotes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	notes, err := h.services.Note.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllNotesResponse{
		Data: notes,
	})
}