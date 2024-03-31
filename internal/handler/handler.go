package handler

import (
	"encoding/json"
	"net/http"

	_ "github.com/Windmill787-golang/junior-test/docs"
	"github.com/Windmill787-golang/junior-test/internal/service"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

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
		r.Post("/", h.CreateBook)
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

func (h *Handler) RespondWithData(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(data)
	if err != nil {
		//TODO: log
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(bytes)
}

func (h *Handler) RespondWithMessage(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := SuccessResponse{Message: message}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		//TODO: log
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(jsonResponse)
}

func (h *Handler) RespondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := ErrorResponse{Error: message}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		//TODO: log
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(jsonResponse)
}
