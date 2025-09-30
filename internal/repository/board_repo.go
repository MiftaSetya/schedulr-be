package repository

import (
	"schedulr/internal/models"

	"gorm.io/gorm"
)

type BoardRepository struct {
	DB *gorm.DB
}

func (r *BoardRepository) CreateBoard(board *models.Board) error {
	return r.DB.Create(board).Error
}

func (r *BoardRepository) GetAllBoards() ([]models.Board, error) {
	var boards []models.Board
	err := r.DB.Find(&boards).Error
	return boards, err
}

func (r *BoardRepository) GetBoardById(id uint) (*models.Board, error) {
	board := &models.Board{}
	err := r.DB.Preload("Lists.Tasks").First(board, id).Error
	return board, err
}

func (r *BoardRepository) GetBoardsByOwner(ownerId uint) ([]models.Board, error) {
	var boards []models.Board
	err := r.DB.Where("owner_id = ?", ownerId).Find(&boards).Error
	return boards, err
}

func (r *BoardRepository) GetBoardsWithListsAndTasks(id uint) (*models.Board, error) {
	board := &models.Board{}
	err := r.DB.
		Preload("Lists", func(db *gorm.DB) *gorm.DB {
			return db.Order("position ASC")
		}).
		Preload("Lists.Tasks", func(db *gorm.DB) *gorm.DB {
			return db.Order("position ASC")
		}).
		First(board, id).Error
	return board, err
}

func (r *BoardRepository) UpdateBoard(board *models.Board) error {
	return r.DB.Save(board).Error
}

func (r *BoardRepository) DeleteBoard(id uint) error {
	return r.DB.Delete(&models.Board{}, id).Error
}
