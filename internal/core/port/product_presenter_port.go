package port

import (
	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/dto"
)

type ProductPresenter interface {
	ToResponse(product *entity.Product) dto.ProductResponse
	ToPaginatedResponse(products []*entity.Product, total int64, page, limit int) dto.PaginatedResponse
}
