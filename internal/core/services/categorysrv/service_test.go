package categorysrv

import (
	"github.com/Sokol111/category-service/internal/core/domain"
	mocks "github.com/Sokol111/category-service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCategoryService_GetById(t *testing.T) {
	repository := mocks.NewMockCategoryRepository(t)
	service := NewCategoryService(repository)
	expected := randomCategory()
	repository.
		EXPECT().
		GetById(expected.ID).
		Return(expected, nil)
	found, err := service.GetById(expected.ID)
	assert.Nil(t, err)
	assert.Equal(t, expected, found)
}

func randomCategory() domain.Category {
	return domain.Category{
		ID:               uuid.NewString(),
		Version:          1,
		Enabled:          true,
		Name:             uuid.NewString(),
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}
}
