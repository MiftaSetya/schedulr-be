package routes

import (
	"schedulr/internal/api/handler"
	"schedulr/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.RouterGroup, h *handler.ProfileHandler) {
	profile := r.Group("/")
	profile.Use(middleware.AuthMiddleware())
	{
		profile.GET("/me", h.Me)
	}
}