package routes

import (
	"schedulr/internal/api/handler"
	"schedulr/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func BoardRoutes(r *gin.RouterGroup, h *handler.BoardHandler) {
	board := r.Group("/")
	board.Use(middleware.AuthMiddleware())
	{
		board.POST("/board", h.CreateBoard)
		board.GET("/board", h.GetBoardsByOwner)
	}
}