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

func (r *CatMemrepo) GetByName(name string) (found model.Category, err error) {
	r.categories.Range(func(_ any, c any) bool {
		category := c.(model.Category)
		if category.Name == name {
			found = category
			return false
		}
		return true
	})
	if found.ID == "" {
		err = fmt.Errorf("failed to find category by name [%v]\n", name)
	}
	return found, err
}

func (r *CatMemrepo) Create(category model.Category) (model.Category, error) {
	category.ID = uuid.NewString()
	category.Version = 1
	category.CreatedDate = time.Now().UTC()
	category.LastModifiedDate = category.CreatedDate
	r.categories.Store(category.ID, category)
	return category, nil
}

func (r *CatMemrepo) Update(category model.Category) (found model.Category, err error) {
	if c, ok := r.categories.Load(category.ID); ok {
		found = c.(model.Category)
		if found.Version != category.Version {
			return model.Category{}, errors.New("failed to update category because of different versions")
		}
		found.Version++
		found.Name = category.Name
		found.LastModifiedDate = time.Now().UTC()
		r.categories.Store(found.ID, found)
	} else {
		err = fmt.Errorf("failed to update category because couldn't find it by id [%v]\n", category.ID)
	}
	return found, err
}

func (r *CatMemrepo) GetCategories() (s []model.Category, err error) {
	r.categories.Range(func(_ any, c any) bool {
		s = append(s, c.(model.Category))
		return true
	})

	return s, err
}
