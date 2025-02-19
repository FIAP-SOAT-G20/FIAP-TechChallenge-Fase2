package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateProductUseCase struct {
	gateway port.ProductGateway
}

// NewUpdateProductUseCase creates a new UpdateProductUseCase
func NewUpdateProductUseCase(gateway port.ProductGateway) port.UpdateProductUseCase {
	return &updateProductUseCase{
		gateway: gateway,
	}
}

// Execute updates a product
func (uc *updateProductUseCase) Execute(ctx context.Context, input dto.UpdateProductInput) (*entity.Product, error) {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if product == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	product.Update(input.Name, input.Description, input.Price, input.CategoryID)

	if err := uc.gateway.Update(ctx, product); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return product, nil
}
