package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteProductUseCase struct {
	gateway port.ProductGateway
}

// NewDeleteProductUseCase creates a new DeleteProductUseCase
func NewDeleteProductUseCase(gateway port.ProductGateway) port.DeleteProductUseCase {
	return &deleteProductUseCase{gateway}
}

// Execute deletes a product
func (uc *deleteProductUseCase) Execute(ctx context.Context, input dto.DeleteProductInput) (*entity.Product, error) {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if product == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return product, nil
}
