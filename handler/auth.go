package handler

import (
	"net/http"

	"github.com/Windmill787-golang/junior-test/entities"
	"github.com/gin-gonic/gin"
)

// @Summary Sign up user
// @Tags auth
// @Description sign up user
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body entities.User true "User info"
// @Success 200 {object} string "Sign up successful"
// @Failure 400 {object} string "Validation error"
// @Failure 500 {object} string "Server error"
// @Router /auth/sign-up [post]
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

// @Summary Sign in user
// @Tags auth
// @Description sign in user
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body entities.User true "User info"
// @Success 200 {object} string "Sign in successful"
// @Failure 400 {object} string "Validation error"
// @Failure 500 {object} string "Server error"
// @Router /auth/sign-in [post]
func (h *Handler) SingIn(c *gin.Context) {
	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Validation error " + err.Error()})
		return
	}

	token, err := h.service.GenerateToken(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server error " + err.Error()})
		return
	}
	if token == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User with these credentials not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Sign in successful", "token": token})
}
