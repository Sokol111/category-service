package service

import (
	"github.com/Sokol111/category-service/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCategoryService_GetById(t *testing.T) {
	repository := mocks.NewMockCategoryRepository(t)
	service := NewCatService(repository)
	expected := randomCategory()
	repository.
		EXPECT().
		GetById(expected.ID).
		Return(expected, nil)
	found, err := service.GetById(expected.ID)
	assert.Nil(t, err)
	assert.Equal(t, expected, found)
}

func randomCategory() model.Category {
	return model.Category{
		ID:               uuid.NewString(),
		Version:          1,
		Enabled:          true,
		Name:             uuid.NewString(),
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}
}
