package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Windmill787-golang/junior-test/internal/entities"
	"github.com/gin-gonic/gin"
)

// GetBooks      godoc
// @Summary      Books list
// @Description  Get books list
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array} entities.Book
// @Failure      500  {object} string "Server error"
// @Router       /books [get]
//func (h *Handler) GetBooks(c *gin.Context) {
//	books, err := h.service.GetBooks()
//	if err != nil {
//		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
//		return
//	}
//
//	c.IndentedJSON(http.StatusOK, books)
//}

// GetBooks      godoc
// @Summary      Books list
// @Description  Get books list
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array} entities.Book
// @Failure      500  {object} string "Server error"
// @Router       /books [get]
func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		//TODO: log
		return
	}

	resp, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		//TODO: log
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		//TODO: log
		return
	}
}

// GetBook       godoc
// @Summary      Book view
// @Description  Get book info
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object} entities.Book
// @Failure      400  {object} string "ID is not provided"
// @Failure      404  {object} string "Book not found"
// @Failure      500  {object} string "Server error"
// @Router       /book/{id} [get]
func (h *Handler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is not provided"})
		return
	}

	book, err := h.service.GetBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	if book == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// CreateBook    godoc
// @Summary      Book create
// @Description  Create new book
// @Tags         books
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        request body entities.Book true "book info"
// @Success      200  {object} string "Book created"
// @Failure      400  {object} string "Validation error"
// @Failure      404  {object} string "Book not found"
// @Failure      500  {object} string "Server error"
// @Router       /book [post]
func (h *Handler) CreateBook(c *gin.Context) {
	var book entities.Book

	if err := c.BindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Validation error " + err.Error()})
		return
	}
	book.UserId = getUserId(c)

	id, err := h.service.CreateBook(book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book created", "id": id})
}

func (h *Handler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is not provided"})
		return
	}

	exist, err := h.service.GetBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	if exist == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if exist.UserId != getUserId(c) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You dont have access to edit this book"})
		return
	}

	var book entities.Book
	if err := c.BindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Validation error " + err.Error()})
		return
	}
	book.ID = id

	if err := h.service.UpdateBook(book); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book updated"})
}

func (h *Handler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is not provided"})
		return
	}

	exist, err := h.service.GetBook(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	if exist.UserId != getUserId(c) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You dont have access to delete this book"})
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func (h *Handler) GetUserId(c *gin.Context) {
	id, _ := c.Get("userId")
	c.IndentedJSON(http.StatusOK, gin.H{"userId": id})
}

func getUserId(c *gin.Context) int {
	userId, exist := c.Get("userId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization needed"})
		return 0
	}

	return userId.(int)
}
