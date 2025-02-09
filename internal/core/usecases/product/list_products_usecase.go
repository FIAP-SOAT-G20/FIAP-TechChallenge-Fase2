package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/dto"
	"tech-challenge-2-app-example/internal/core/port"
)

type listProductsUseCase struct {
	repository port.ProductRepository
	presenter  port.ProductPresenter
}

func NewListProductsUseCase(repo port.ProductRepository, presenter port.ProductPresenter) port.ListProductsUseCase {
	return &listProductsUseCase{
		repository: repo,
		presenter:  presenter,
	}
}

func (uc *listProductsUseCase) Execute(ctx context.Context, req dto.ProductListRequest) (*dto.PaginatedResponse, error) {
	// Validações adicionais podem ser feitas aqui
	if req.Page < 1 {
		return nil, errors.NewInvalidInputError("Página deve ser maior que zero")
	}

	if req.Limit < 1 || req.Limit > 100 {
		return nil, errors.NewInvalidInputError("Limite deve estar entre 1 e 100")
	}

	products, total, err := uc.repository.FindAll(ctx, req.Name, req.CategoryID, req.Page, req.Limit)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if len(products) == 0 {
		return nil, errors.NewNotFoundError("Nenhum produto encontrado")
	}

	response := uc.presenter.ToPaginatedResponse(products, total, req.Page, req.Limit)
	return &response, nil
}
