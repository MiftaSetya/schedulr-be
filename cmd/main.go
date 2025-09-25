package main

import (
	"log"
	"os"
	"schedulr/internal/api/handler"
	"schedulr/internal/api/routes"
	"schedulr/internal/config"
	"schedulr/internal/models"
	"schedulr/internal/repository"
	"schedulr/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	jwtSecret := os.Getenv("JWT_SECRET")

	userRepo := &repository.UserRepository{DB: db}
	authService := &service.AuthService{Repo: userRepo}
	authHandler := &handler.AuthHandler{Service: authService, JwtSecret: jwtSecret}

	api := r.Group("/api")
	{
		routes.AuthRoutes(api, authHandler)
	}

	r.Run(":8080")
}