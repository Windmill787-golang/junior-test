package handler

import (
	"github.com/Windmill787-golang/junior-test/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//create routes for book
	router := gin.New()

	router.GET("/book/:id", h.GetBook)

	return router
}
