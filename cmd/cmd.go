package cmd

import (
	"log"
	"os"
	"schedulr/internal/api/handler"
	"schedulr/internal/api/routes"
	"schedulr/internal/config"
	"schedulr/internal/models"
	"schedulr/internal/repository"
	"schedulr/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Exec() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.List{},
		&models.Task{},
	)

	jwtSecret := os.Getenv("JWT_SECRET")

	// Init Repos
	userRepo := &repository.UserRepository{DB: db}
	boardRepo := &repository.BoardRepository{DB: db}

	// Init Service
	authService := &service.AuthService{Repo: userRepo}
	boardService := &service.BoardService{Repo: boardRepo}

	// Init Handler
	authHandler := &handler.AuthHandler{Service: authService, JwtSecret: jwtSecret}
	profileHandler := &handler.ProfileHandler{UserRepo: userRepo}
	boardHandler := &handler.BoardHandler{Service: boardService}

	api := r.Group("/api")
	{
		routes.AuthRoutes(api, authHandler)
		routes.ProfileRoutes(api, profileHandler)
		routes.BoardRoutes(api, boardHandler)
	}

	r.Run(":8080")
}
