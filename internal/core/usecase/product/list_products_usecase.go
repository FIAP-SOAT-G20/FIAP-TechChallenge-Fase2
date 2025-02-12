package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listProductsUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

// NewListProductsUseCase creates a new ListProductsUseCase
func NewListProductsUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.ListProductsUseCase {
	return &listProductsUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute lists all products
func (uc *listProductsUseCase) Execute(ctx context.Context, input dto.ListProductsInput) error {
	products, total, err := uc.gateway.FindAll(ctx, input.Name, input.CategoryID, input.Page, input.Limit)
	if err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.ProductPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: products,
	})
	return nil
}
