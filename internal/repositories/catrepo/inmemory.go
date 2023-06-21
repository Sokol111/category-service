package catrepo

import (
	"errors"
	"fmt"
	"github.com/Sokol111/category-service/internal/core/domain"
	"github.com/Sokol111/category-service/internal/core/ports"
	"github.com/google/uuid"
	"sync"
	"time"
)

type memrepo struct {
	categories map[string]domain.Category
	mu         sync.RWMutex
}

func NewCategoryRepository() ports.CategoryRepository {
	return &memrepo{categories: map[string]domain.Category{}}
}

func (r *memrepo) GetById(id string) (domain.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if category, ok := r.categories[id]; ok {
		return category, nil
	}
	return domain.Category{}, fmt.Errorf("failed to find category by id [%v]\n", id)
}

func (r *memrepo) GetByName(name string) (domain.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, v := range r.categories {
		if v.Name == name {
			return v, nil
		}
	}
	return domain.Category{}, fmt.Errorf("failed to find category by name [%v]\n", name)
}

func (r *memrepo) Create(category domain.Category) (domain.Category, error) {
	category.ID = uuid.NewString()
	category.Version = 1
	category.CreatedDate = time.Now().UTC()
	category.LastModifiedDate = category.CreatedDate
	r.mu.Lock()
	r.categories[category.ID] = category
	r.mu.Unlock()
	return category, nil
}

func (r *memrepo) Update(category domain.Category) (domain.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if found, ok := r.categories[category.ID]; ok {
		if found.Version != category.Version {
			return domain.Category{}, errors.New("failed to update category because of different versions")
		}
		found.Version++
		found.Name = category.Name
		found.LastModifiedDate = time.Now().UTC()
		r.categories[found.ID] = found
	} else {
		return domain.Category{}, fmt.Errorf("failed to update category because couldn't find it by id [%v]\n", category.ID)
	}
	return category, nil
}

func (r *memrepo) GetCategories() ([]domain.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s := make([]domain.Category, 0, len(r.categories))
	for _, v := range r.categories {
		s = append(s, v)
	}
	return s, nil
}
