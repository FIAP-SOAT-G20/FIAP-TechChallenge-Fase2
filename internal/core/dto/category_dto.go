package dto

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"Foods"`
}

func (r CreateCategoryRequest) ToEntity() *entity.Category {
	return &entity.Category{
		Name: r.Name,
	}
}

type GetCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

type ListCategoriesRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"omitempty,required" example:"Beverages"`
}

func (r UpdateCategoryRequest) ToEntity(id uint64) *entity.Category {
	return &entity.Category{
		ID:   id,
		Name: r.Name,
	}
}

type DeleteCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

type CategoryResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCategoryResponse(category *entity.Category) *CategoryResponse {
	if category == nil {
		return nil
	}

	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

type CategoriesPaginatedResponse struct {
	presenter.JsonPagination
	Categories []CategoryResponse `json:"categories"`
}

func NewCategoriesPaginatedResponse(categories []entity.Category, total int64, page int, limit int) *CategoriesPaginatedResponse {
	categoryResponses := make([]CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponse := NewCategoryResponse(&category)
		if categoryResponse != nil {
			categoryResponses = append(categoryResponses, *categoryResponse)
		}
	}

	return &CategoriesPaginatedResponse{
		JsonPagination: presenter.JsonPagination{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Categories: categoryResponses,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
