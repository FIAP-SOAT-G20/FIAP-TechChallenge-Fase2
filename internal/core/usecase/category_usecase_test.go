package usecase

import (
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupCategoryTest() (*CategoryUsecase, *mocks.CategoryDatasourceMock) {
	mockDatasource := new(mocks.CategoryDatasourceMock)
	usecase := NewCategoryUsecase(mockDatasource)
	return usecase, mockDatasource
}

func TestCategoryUsecase_Create(t *testing.T) {
	usecase, mockDatasource := setupCategoryTest()

	t.Run("success", func(t *testing.T) {
		category := &entity.Category{
			Name: "Test Category",
		}

		mockDatasource.On("Insert", mock.AnythingOfType("*entity.Category")).Return(nil)

		err := usecase.Create(category)

		assert.NoError(t, err)
		mockDatasource.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		category := &entity.Category{
			Name: "Test Category",
		}

		mockDatasource.On("Insert", mock.AnythingOfType("*entity.Category")).Return(entity.ErrNotFound)

		err := usecase.Create(category)

		assert.Error(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		mockDatasource.AssertExpectations(t)
	})
}

func TestCategoryUsecase_GetByID(t *testing.T) {
	usecase, mockDatasource := setupCategoryTest()

	t.Run("success", func(t *testing.T) {
		expectedCategory := &entity.Category{
			ID:   1,
			Name: "Test Category",
		}

		mockDatasource.On("GetByID", uint64(1)).Return(expectedCategory, nil)

		category, err := usecase.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, category)
		mockDatasource.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockDatasource.On("GetByID", uint64(999)).Return(nil, entity.ErrNotFound)

		category, err := usecase.GetByID(999)

		assert.Error(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, category)
		mockDatasource.AssertExpectations(t)
	})
}

func TestCategoryUsecase_List(t *testing.T) {
	usecase, mockDatasource := setupCategoryTest()

	t.Run("success", func(t *testing.T) {
		expectedCategories := []entity.Category{
			{ID: 1, Name: "Category 1"},
			{ID: 2, Name: "Category 2"},
		}
		expectedTotal := int64(2)

		mockDatasource.On("GetAll", "test", 1, 10).Return(expectedCategories, expectedTotal, nil)

		categories, total, err := usecase.List("test", 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, categories)
		assert.Equal(t, expectedTotal, total)
		mockDatasource.AssertExpectations(t)
	})

	t.Run("empty result", func(t *testing.T) {
		mockDatasource.On("GetAll", "nonexistent", 1, 10).Return([]entity.Category{}, int64(0), nil)

		categories, total, err := usecase.List("nonexistent", 1, 10)

		assert.NoError(t, err)
		assert.Empty(t, categories)
		assert.Equal(t, int64(0), total)
		mockDatasource.AssertExpectations(t)
	})
}

func TestCategoryUsecase_Update(t *testing.T) {
	usecase, mockDatasource := setupCategoryTest()

	t.Run("success", func(t *testing.T) {
		category := &entity.Category{
			ID:   1,
			Name: "Updated Category",
		}

		mockDatasource.On("GetByID", uint64(1)).Return(category, nil)
		mockDatasource.On("Update", mock.AnythingOfType("*entity.Category")).Return(nil)

		err := usecase.Update(category)

		assert.NoError(t, err)
		mockDatasource.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		category := &entity.Category{
			ID:   999,
			Name: "Updated Category",
		}

		mockDatasource.On("GetByID", uint64(999)).Return(nil, entity.ErrNotFound)

		err := usecase.Update(category)

		assert.Error(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		mockDatasource.AssertExpectations(t)
	})
}

func TestCategoryUsecase_Delete(t *testing.T) {
	usecase, mockDatasource := setupCategoryTest()

	t.Run("success", func(t *testing.T) {
		mockDatasource.On("GetByID", uint64(1)).Return(&entity.Category{ID: 1}, nil)
		mockDatasource.On("Delete", uint64(1)).Return(nil)

		err := usecase.Delete(1)

		assert.NoError(t, err)
		mockDatasource.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockDatasource.On("GetByID", uint64(999)).Return(nil, entity.ErrNotFound)

		err := usecase.Delete(999)

		assert.Error(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		mockDatasource.AssertExpectations(t)
	})
} 