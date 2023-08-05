package handler

import (
	"net/http"
	"strconv"

	"github.com/Windmill787-golang/junior-test/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBook(c *gin.Context) {
	//get id from context
	//json indent book
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
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

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.service.GetBooks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func (h *Handler) CreateBook(c *gin.Context) {
	var book entities.Book

	if err := c.BindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Validation error " + err.Error()})
		return
	}

	id, err := h.service.CreateBook(book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book created", "id": id})
}
