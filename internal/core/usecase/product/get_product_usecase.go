package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getProductUseCase struct {
	gateway port.ProductGateway
}

// NewGetProductUseCase creates a new GetProductUseCase
func NewGetProductUseCase(gateway port.ProductGateway) port.GetProductUseCase {
	return &getProductUseCase{gateway}
}

// Execute gets a product
func (uc *getProductUseCase) Execute(ctx context.Context, input dto.GetProductInput) (*entity.Product, error) {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if product == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return product, nil
}
