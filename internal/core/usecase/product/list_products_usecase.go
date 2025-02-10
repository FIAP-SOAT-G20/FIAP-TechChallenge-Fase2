package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
)

type listProductsUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

func NewListProductsUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.ListProductsUseCase {
	return &listProductsUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

func (uc *listProductsUseCase) Execute(ctx context.Context, input usecase.ListProductsInput) (*usecase.ListProductPaginatedOutput, error) {
	if input.Page < 1 {
		return nil, domain.NewInvalidInputError("página deve ser maior que zero")
	}

	if input.Limit < 1 || input.Limit > 100 {
		return nil, domain.NewInvalidInputError("limite deve estar entre 1 e 100")
	}

	products, total, err := uc.gateway.FindAll(ctx, input.Name, input.CategoryID, input.Page, input.Limit)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	return uc.presenter.ToPaginatedOutput(products, total, input.Page, input.Limit), nil
}
