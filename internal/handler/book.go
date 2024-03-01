package handler

import (
	"net/http"
	"strconv"

	"github.com/Windmill787-golang/junior-test/internal/entities"
	"github.com/gin-gonic/gin"
)

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
