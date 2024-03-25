package handler

import (
	_ "github.com/Windmill787-golang/junior-test/docs"
	"github.com/Windmill787-golang/junior-test/internal/service"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Route("/books", func(r chi.Router) {
		r.Get("/{id:[0-9]+}", h.GetBook)
		r.Get("/", h.GetBooks)
	})

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

	return r
}
