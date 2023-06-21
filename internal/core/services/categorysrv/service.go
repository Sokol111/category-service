package categorysrv

import (
	"github.com/Sokol111/category-service/internal/core/domain"
	"github.com/Sokol111/category-service/internal/core/ports"
)

type service struct {
	repository ports.CategoryRepository
}

func NewCategoryService(repository ports.CategoryRepository) ports.CategoryService {
	return &service{repository}
}

func (s *service) GetById(id string) (domain.Category, error) {
	return s.repository.GetById(id)
}

func (s *service) GetByName(name string) (domain.Category, error) {
	return s.repository.GetByName(name)
}

func (s *service) Create(category domain.Category) (domain.Category, error) {
	return s.repository.Create(category)
}

func (s *service) Update(category domain.Category) (domain.Category, error) {
	return s.repository.Update(category)
}

func (s *service) GetCategories() ([]domain.Category, error) {
	return s.repository.GetCategories()
}
