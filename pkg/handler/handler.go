package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app/pkg/service"

	_ "github.com/zhashkevych/todo-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	return router
}
