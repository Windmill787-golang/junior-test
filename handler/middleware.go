package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header not provided"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is invalid"})
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization error " + err.Error()})
		return
	}

	if userId == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is invalid"})
		return
	}

	c.Set("userId", userId)
}
