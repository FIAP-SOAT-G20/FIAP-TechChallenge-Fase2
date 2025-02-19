package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createProductUseCase struct {
	gateway port.ProductGateway
}

// NewCreateProductUseCase creates a new CreateProductUseCase
func NewCreateProductUseCase(gateway port.ProductGateway) port.CreateProductUseCase {
	return &createProductUseCase{gateway}
}

// Execute creates a new product
func (uc *createProductUseCase) Execute(ctx context.Context, input dto.CreateProductInput) (*entity.Product, error) {
	product := entity.NewProduct(input.Name, input.Description, input.Price, input.CategoryID)

	if err := uc.gateway.Create(ctx, product); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return product, nil
}
