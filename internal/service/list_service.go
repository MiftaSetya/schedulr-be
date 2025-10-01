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

func (s *ListService) UpdateListPosition(input *dto.UpdateListPositionInput) error {
	return s.Repo.UpdateListPosition(input.ID, input.Position)
}

func (s *ListService) UpdateListsOrder(input *dto.UpdateListsOrderInput) error {
	var lists []models.List
	for _, l := range input.Lists {
		list := models.List{
			Position: l.Position,
		}
		list.ID = l.ID 

		lists = append(lists, list)
	}
	return s.Repo.UpdateListsOrder(lists)
}


func (s *ListService) DeleteList(id uint) error {
	return s.Repo.DeleteList(id)
}
