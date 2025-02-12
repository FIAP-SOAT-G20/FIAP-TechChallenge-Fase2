package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

// NewUpdateProductUseCase creates a new UpdateProductUseCase
func NewUpdateProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.UpdateProductUseCase {
	return &updateProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute updates a product
func (uc *updateProductUseCase) Execute(ctx context.Context, input dto.UpdateProductInput) error {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if product == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	product.Update(input.Name, input.Description, input.Price, input.CategoryID)

	if err := uc.gateway.Update(ctx, product); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.ProductPresenterInput{
		Writer: input.Writer,
		Result: product,
	})
	return nil
}
