package port

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
)

type ProductPresenter interface {
	ToOutput(product *entity.Product) *usecase.ProductOutput
	ToPaginatedOutput(products []*entity.Product, total int64, page, limit int) *usecase.ListProductPaginatedOutput
}
