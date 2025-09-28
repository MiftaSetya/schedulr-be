package handler

import (
	"net/http"
	"schedulr/internal/dto"
	"schedulr/internal/repository"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	UserRepo *repository.UserRepository
}

func (h *ProfileHandler) Me(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	user, err := h.UserRepo.GetUserById(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	userResp := dto.UserResponse {
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, gin.H{"user": userResp})
}