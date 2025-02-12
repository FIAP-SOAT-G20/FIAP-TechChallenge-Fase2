package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

// NewDeleteProductUseCase creates a new DeleteProductUseCase
func NewDeleteProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.DeleteProductUseCase {
	return &deleteProductUseCase{gateway, presenter}
}

// Execute deletes a product
func (uc *deleteProductUseCase) Execute(ctx context.Context, input dto.DeleteProductInput) error {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if product == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.ProductPresenterInput{
		Writer: input.Writer,
		Result: product,
	})

	return nil
}
