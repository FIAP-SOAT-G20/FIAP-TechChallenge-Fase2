package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listProductsUseCase struct {
	gateway port.ProductGateway
}

// NewListProductsUseCase creates a new ListProductsUseCase
func NewListProductsUseCase(gateway port.ProductGateway) port.ListProductsUseCase {
	return &listProductsUseCase{gateway}
}

// Execute lists all products
func (uc *listProductsUseCase) Execute(ctx context.Context, input dto.ListProductsInput) ([]*entity.Product, int64, error) {
	products, total, err := uc.gateway.FindAll(ctx, input.Name, input.CategoryID, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return products, total, nil
}
