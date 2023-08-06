package handler

import (
	"net/http"

	"github.com/Windmill787-golang/junior-test/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SingUp(c *gin.Context) {
	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Validation error " + err.Error()})
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Sign up successful", "id": id})
}
