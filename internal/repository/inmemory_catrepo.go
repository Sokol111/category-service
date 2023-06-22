package repository

import (
	"errors"
	"fmt"
	"github.com/Sokol111/category-service/internal/model"
	"github.com/google/uuid"
	"sync"
	"time"
)

type CatMemrepo struct {
	categories sync.Map
}

func NewCatInmemoryRepo() *CatMemrepo {
	return &CatMemrepo{categories: sync.Map{}}
}

func (r *CatMemrepo) GetById(id string) (model.Category, error) {
	if category, ok := r.categories.Load(id); ok {
		return category.(model.Category), nil
	}

	return model.Category{}, fmt.Errorf("failed to find category by id [%v]\n", id)
}

func (r *CatMemrepo) GetByName(name string) (model.Category, error) {
	for _, v := range r.categories {
		if v.Name == name {
			return v, nil
		}
	}
	return model.Category{}, fmt.Errorf("failed to find category by name [%v]\n", name)
}

func (r *CatMemrepo) Create(category model.Category) (model.Category, error) {
	category.ID = uuid.NewString()
	category.Version = 1
	category.CreatedDate = time.Now().UTC()
	category.LastModifiedDate = category.CreatedDate
	r.categories[category.ID] = category

	return category, nil
}

func (r *CatMemrepo) Update(category model.Category) (model.Category, error) {
	if found, ok := r.categories[category.ID]; ok {
		if found.Version != category.Version {
			return model.Category{}, errors.New("failed to update category because of different versions")
		}
		found.Version++
		found.Name = category.Name
		found.LastModifiedDate = time.Now().UTC()
		r.categories[found.ID] = found
	} else {
		return model.Category{}, fmt.Errorf("failed to update category because couldn't find it by id [%v]\n", category.ID)
	}
	return category, nil
}

func (r *CatMemrepo) GetCategories() ([]model.Category, error) {
	s := make([]model.Category, 0, len(r.categories))
	for _, v := range r.categories {
		s = append(s, v)
	}
	return s, nil
}
