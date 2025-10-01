package service

import (
	"schedulr/internal/dto"
	"schedulr/internal/models"
	"schedulr/internal/repository"
)

type BoardService struct {
	Repo *repository.BoardRepository
}

func (s *BoardService) CreateBoard(input dto.CreateBoardInput) (*models.Board, error) {
	board := &models.Board{
		Name: input.Name,
		OwnerID: input.OwnerID,
	}

	if err := s.Repo.CreateBoard(board); err != nil {
		return nil, err
	}

	return board, nil
}

func (s *BoardService) GetBoardById(id uint) (*models.Board, error) {
	return s.Repo.GetBoardById(id)
}

func (s *BoardService) GetBoardsWithListsAndTasks(id uint) (*models.Board, error) {
	return s.Repo.GetBoardsWithListsAndTasks(id)
}

func (s *BoardService) GetBoardsByOwnerId(OwnerID uint) ([]models.Board, error) {
	return s.Repo.GetBoardsByOwner(OwnerID)
}

func (s *BoardService) UpdateBoard(id uint, input dto.CreateBoardInput) (*models.Board, error) {
	board, err := s.Repo.GetBoardById(id)
	if err != nil {
		return nil, err
	}

	board.Name = input.Name
	board.OwnerID = input.OwnerID

	if err := s.Repo.UpdateBoard(board); err != nil {
		return nil, err
	}

	return board, nil
}

func (s *BoardService) DeleteBoard(id uint) error {
	return s.Repo.DeleteBoard(id)
}