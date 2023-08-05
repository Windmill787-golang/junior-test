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

func InitRoutes() *gin.Engine {
	//create routes for book
	return gin.Default()
}
