package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest() (*gin.Engine, *mocks.CategoryUsecaseMock) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUsecase := new(mocks.CategoryUsecaseMock)
	controller := NewCategoryController(mockUsecase)
	
	group := r.Group(controller.GroupRouterPattern())
	controller.Register(group)
	
	return r, mockUsecase
}

func TestCategoryController_CreateCategory(t *testing.T) {
	r, mockUsecase := setupTest()

	t.Run("success", func(t *testing.T) {
		req := dto.CreateCategoryRequest{Name: "Test Category"}
		category := &entity.Category{Name: req.Name}
		
		mockUsecase.On("Create", mock.AnythingOfType("*entity.Category")).Return(nil)

		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/categories/", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("invalid request", func(t *testing.T) {
		req := dto.CreateCategoryRequest{} // Nome vazio
		
		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/categories/", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCategoryController_GetCategory(t *testing.T) {
	r, mockUsecase := setupTest()

	t.Run("success", func(t *testing.T) {
		category := &entity.Category{ID: 1, Name: "Test Category"}
		mockUsecase.On("GetByID", uint64(1)).Return(category, nil)

		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/categories/1", nil)
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUsecase.On("GetByID", uint64(999)).Return(nil, entity.ErrNotFound)

		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/categories/999", nil)
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestCategoryController_ListCategories(t *testing.T) {
	r, mockUsecase := setupTest()

	t.Run("success", func(t *testing.T) {
		categories := []entity.Category{{ID: 1, Name: "Category 1"}}
		mockUsecase.On("List", "", 1, 10).Return(categories, int64(1), nil)

		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/categories/", nil)
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestCategoryController_UpdateCategory(t *testing.T) {
	r, mockUsecase := setupTest()

	t.Run("success", func(t *testing.T) {
		req := dto.UpdateCategoryRequest{Name: "Updated Category"}
		category := &entity.Category{ID: 1, Name: req.Name}
		
		mockUsecase.On("Update", mock.AnythingOfType("*entity.Category")).Return(nil)

		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPut, "/api/v1/categories/1", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestCategoryController_DeleteCategory(t *testing.T) {
	r, mockUsecase := setupTest()

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("Delete", uint64(1)).Return(nil)

		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodDelete, "/api/v1/categories/1", nil)
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusNoContent, w.Code)
		mockUsecase.AssertExpectations(t)
	})
} 