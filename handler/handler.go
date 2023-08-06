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
	router := gin.Default()

	router.GET("/books", h.GetBooks)

	book := router.Group("/book")
	{
		book.POST("/", h.CreateBook)
		book.GET("/:id", h.GetBook)
		book.PUT("/:id", h.UpdateBook)
		book.DELETE("/:id", h.DeleteBook)
	}

	user := router.Group("/auth")
	{
		user.POST("/sign-up", h.SingUp)
	}

	return router
}
