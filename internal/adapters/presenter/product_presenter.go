package presenter

import (
	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/port"
	"tech-challenge-2-app-example/internal/core/usecase"
)

type productPresenter struct{}

func NewProductPresenter() port.ProductPresenter {
	return &productPresenter{}
}

func (p *productPresenter) ToOutput(product *entity.Product) *usecase.ProductOutput {
	return &usecase.ProductOutput{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (p *productPresenter) ToPaginatedOutput(products []*entity.Product, total int64, page, limit int) *usecase.ListProductPaginatedOutput {
	productOutputs := make([]usecase.ProductOutput, len(products))
	for i, product := range products {
		output := p.ToOutput(product)
		productOutputs[i] = *output
	}

	return &usecase.ListProductPaginatedOutput{
		PaginatedOutput: usecase.PaginatedOutput{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Products: productOutputs,
	}
}
