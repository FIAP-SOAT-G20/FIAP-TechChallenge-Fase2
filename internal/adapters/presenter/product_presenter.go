package presenter

import (
	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/dto"
	"tech-challenge-2-app-example/internal/core/port"
)

type productPresenter struct{}

func NewProductPresenter() port.ProductPresenter {
	return &productPresenter{}
}

func (p *productPresenter) ToResponse(product *entity.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (p *productPresenter) ToPaginatedResponse(products []*entity.Product, total int64, page, limit int) dto.PaginatedResponse {
	var responses []dto.ProductResponse
	for _, product := range products {
		responses = append(responses, p.ToResponse(product))
	}

	return dto.PaginatedResponse{
		Total:    total,
		Page:     page,
		Limit:    limit,
		Products: responses,
	}
}
