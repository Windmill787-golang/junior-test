package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Windmill787-golang/junior-test/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"

	_ "github.com/Windmill787-golang/junior-test/docs"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/books", h.GetBooks)
	router.GET("/book/:id", h.GetBook)

	book := router.Group("/book", h.userIdentity)
	{
		book.POST("/", h.CreateBook)
		book.PUT("/:id", h.UpdateBook)
		book.DELETE("/:id", h.DeleteBook)
	}

	user := router.Group("/auth")
	{
		user.POST("/sign-up", h.SingUp)
		user.POST("/sign-in", h.SingIn)
	}

	router.GET("/user-id", h.GetUserId)

	return router
}

func (h *Handler) InitRoutesNew() {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		books, err := h.service.GetBooks()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			//log
			return
		}

		resp, err := json.Marshal(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			//log
			return
		}

		n, err := w.Write(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			//log
			return
		}
		fmt.Println(n)
	})
}
