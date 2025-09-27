package handler

import (
	"net/http"
	"schedulr/internal/dto"
	"schedulr/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	Service   *service.AuthService
	JwtSecret string
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input *dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered",
		"user":    dto.ToUserResponse(user),
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(30 * 24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.JwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		User:  dto.ToUserResponse(user),
		Token: tokenString,
	})
}
