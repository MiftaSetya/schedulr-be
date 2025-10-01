package service

import (
	"schedulr/internal/dto"
	"schedulr/internal/models"
	"schedulr/internal/repository"
)

type ListService struct {
	Repo *repository.ListRepository
}

func (s *ListService) CreateList(input *dto.CreateListInput) (*models.List, error) {
	list := &models.List{
		Name:     input.Name,
		BoardID:  input.BoardID,
		Position: input.Position,
	}

	if err := s.Repo.CreateList(list); err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ListService) GetListByID(id uint) (*models.List, error) {
	return s.Repo.GetListById(id)
}

func (s *ListService) UpdateListName(id uint, input *dto.CreateListInput) (*models.List, error) {
	list, err := s.Repo.GetListById(id)
	if err != nil {
		return nil, err
	}

	list.BoardID = input.BoardID
	list.Name = input.Name

	if err := s.Repo.UpdateList(list); err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ListService) UpdateListPosition(id uint, newPosition int) (*models.List, error) {
	list, err := s.Repo.GetListById(id)
	if err != nil {
		return nil, err
	}

	list.Position = newPosition

	if err := s.Repo.UpdateList(list); err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ListService) UpdateListOrder(lists []*models.List) error {
	for _, list := range lists {
		if err := s.Repo.UpdateList(list); err != nil {
			return err
		}
	}
	return nil
}

func (s *ListService) DeleteList(id uint) error {
	return s.Repo.DeleteList(id)
}
