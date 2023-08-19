package handler

import (
	"net/http"

	"github.com/mrnkslv/kodeProject/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/auth/sign-in", h.signIn)
	router.HandleFunc("/auth/sign-up", h.signUp)

	router.HandleFunc("/api/notes/", h.myMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.getNotes(w, r)
		case http.MethodPost:
			h.createNote(w, r)
		default:
			newErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	}))

	return router
}

func (h *Handler) myMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.userIdentity(w, r)
		next(w, r)
	}

}
