package service

import (
	"errors"
	"schedulr/internal/models"
	"schedulr/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func (s *AuthService) Register(user *models.User) error {
	existing, _ := s.Repo.GetUserByEmail(user.Email)
	if existing.ID != 0 {
		return errors.New("email already resgitered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return s.Repo.CreateUser(user)
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