package port

import (
	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/usecase"
)

type ProductPresenter interface {
	ToOutput(product *entity.Product) *usecase.ProductOutput
	ToPaginatedOutput(products []*entity.Product, total int64, page, limit int) *usecase.ListProductPaginatedOutput
}
