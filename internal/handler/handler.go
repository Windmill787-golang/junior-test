package handler

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	_ "github.com/Windmill787-golang/junior-test/docs"
	"github.com/Windmill787-golang/junior-test/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.DefaultServeMux
	router.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/swagger.json")),
	)

	router.HandleFunc("/books", h.GetBooks)

	//router.GET("/book/:id", h.GetBook)
	//
	//book := router.Group("/book", h.userIdentity)
	//{
	//	book.POST("/", h.CreateBook)
	//	book.PUT("/:id", h.UpdateBook)
	//	book.DELETE("/:id", h.DeleteBook)
	//}
	//
	//user := router.Group("/auth")
	//{
	//	user.POST("/sign-up", h.SingUp)
	//	user.POST("/sign-in", h.SingIn)
	//}
	//
	//router.GET("/user-id", h.GetUserId)

	return router
}
