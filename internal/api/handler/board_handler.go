package handler

import (
	"net/http"
	"schedulr/internal/dto"
	"schedulr/internal/service"

	"github.com/gin-gonic/gin"
)

type BoardHandler struct {
	Service *service.BoardService
}

func (h *BoardHandler) CreateBoard(c *gin.Context) {
	var input dto.CreateBoardInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ 
			"error": "Invalid request body",
			"detail": err.Error(),
		})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	input.OwnerID = userID.(uint)

	board, err := h.Service.CreateBoard(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create board",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Board created successfully",
		"data": board,
	})
}

func (h *BoardHandler) GetBoardsByOwner(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentiaction token missing or invalid",
		})
		return
	}

	boards, err := h.Service.GetBoardsByOwnerId(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve boards",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": boards,
	})
}