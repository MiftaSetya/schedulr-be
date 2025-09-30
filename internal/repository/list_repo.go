package repository

import (
	"schedulr/internal/models"

	"gorm.io/gorm"
)

type ListRepository struct {
	DB *gorm.DB
}

func (r *ListRepository) CreateList(list *models.List) error {
	return r.DB.Create(list).Error
}

func (r *ListRepository) GetAllLists() ([]models.List, error) {
	var lists []models.List
	err := r.DB.Find(&lists).Error
	return lists, err
}

func (r *ListRepository) GetListById(id uint) (*models.List, error) {
	list := &models.List{}
	err := r.DB.Preload("Tasks").First(list, id).Error
	return list, err
}

func (r *ListRepository) UpdateList(list *models.List) error {
	return r.DB.Save(list).Error
}

func (r *ListRepository) UpdateListPosition(listID uint, newPos int) error {
    return r.DB.Model(&models.List{}).
        Where("id = ?", listID).
        Update("position", newPos).Error
}

func (r *ListRepository) DeleteList(id uint) error {
	return r.DB.Delete(&models.List{}, id).Error
}
