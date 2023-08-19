package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get(authorizationHeader)
	if header == "" {
		newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	r.Header.Set(userCtx, fmt.Sprint(userId))
}

func getUserId(w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.Header.Get(userCtx)
	if id == "" {
		newErrorResponse(w, http.StatusInternalServerError, "user id not found.")
		return 0, errors.New("user id not found")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "invalid type of user id")
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}
