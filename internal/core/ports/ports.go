package ports

import "github.com/Sokol111/category-service/internal/core/domain"

type CategoryService interface {
	GetById(id string) (domain.Category, error)

	GetByName(name string) (domain.Category, error)

	Create(category domain.Category) (domain.Category, error)

	Update(category domain.Category) (domain.Category, error)

	GetCategories() ([]domain.Category, error)
}

type CategoryRepository interface {
	GetById(id string) (domain.Category, error)

	GetByName(name string) (domain.Category, error)

	Create(category domain.Category) (domain.Category, error)

	Update(category domain.Category) (domain.Category, error)

	GetCategories() ([]domain.Category, error)
}
