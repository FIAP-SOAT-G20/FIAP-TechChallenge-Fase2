package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/port"
	"tech-challenge-2-app-example/internal/core/usecase"
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
		return nil, errors.NewInvalidInputError("página deve ser maior que zero")
	}

	if input.Limit < 1 || input.Limit > 100 {
		return nil, errors.NewInvalidInputError("limite deve estar entre 1 e 100")
	}

	products, total, err := uc.gateway.FindAll(ctx, input.Name, input.CategoryID, input.Page, input.Limit)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return uc.presenter.ToPaginatedOutput(products, total, input.Page, input.Limit), nil
}
