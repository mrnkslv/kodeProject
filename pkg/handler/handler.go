package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrnkslv/kodeProject/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		notes := api.Group("/notes")
		{
			notes.POST("/", h.createNote)
			notes.GET("/", h.getNotes)
		}
	}
	return router
}
