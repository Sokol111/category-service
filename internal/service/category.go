package service

import (
	"github.com/Sokol111/category-service/internal/model"
)

type CatService struct {
	repository CategoryRepository
}

type CategoryService interface {
	GetById(id string) (model.Category, error)

	GetByName(name string) (model.Category, error)

	Create(category model.Category) (model.Category, error)

	Update(category model.Category) (model.Category, error)

	GetCategories() ([]model.Category, error)
}

type CategoryRepository interface {
	CategoryService
}

func NewCatService(repository CategoryRepository) *CatService {
	return &CatService{repository}
}

func (s *CatService) GetById(id string) (model.Category, error) {
	return s.repository.GetById(id)
}

func (s *CatService) GetByName(name string) (model.Category, error) {
	return s.repository.GetByName(name)
}

func (s *CatService) Create(category model.Category) (model.Category, error) {
	return s.repository.Create(category)
}

func (s *CatService) Update(category model.Category) (model.Category, error) {
	return s.repository.Update(category)
}

func (s *CatService) GetCategories() ([]model.Category, error) {
	return s.repository.GetCategories()
}
