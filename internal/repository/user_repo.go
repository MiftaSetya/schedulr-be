package repository

import (
	"schedulr/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.DB.Where("email = ?", email).First(user).Error
	return user, err
}

func (r *UserRepository) GetUserById(id uint) (*models.User, error) {
	user := &models.User{}
	err := r.DB.First(user, id).Error
	return user, err
}