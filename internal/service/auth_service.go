package service

import (
	"errors"
	"schedulr/internal/dto"
	"schedulr/internal/models"
	"schedulr/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func (s *AuthService) Register(input *dto.RegisterInput) (*models.User, error) {
	existing, _ := s.Repo.GetUserByEmail(input.Email)
	if existing.ID != 0 {
		return nil, errors.New("email already resgitered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}

	if err := s.Repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
