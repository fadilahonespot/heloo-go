package service

import (
	"errors"

	"heloo-go/internal/domain"
	"heloo-go/internal/repository"
)

type ItemService interface {
	List() ([]domain.Item, error)
	Get(id string) (domain.Item, error)
	Create(in domain.Item) (domain.Item, error)
	Update(in domain.Item) (domain.Item, error)
	Delete(id string) error
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService { return &itemService{repo: repo} }

func (s *itemService) List() ([]domain.Item, error) { return s.repo.List() }

func (s *itemService) Get(id string) (domain.Item, error) { return s.repo.GetByID(id) }

func (s *itemService) Create(in domain.Item) (domain.Item, error) {
	if in.Title == "" {
		return domain.Item{}, errors.New("title is required")
	}
	return s.repo.Create(in)
}

func (s *itemService) Update(in domain.Item) (domain.Item, error) {
	if in.ID == 0 {
		return domain.Item{}, errors.New("id is required")
	}
	if in.Title == "" {
		return domain.Item{}, errors.New("title is required")
	}
	return s.repo.Update(in)
}

func (s *itemService) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	affected, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("not found")
	}
	return nil
}
