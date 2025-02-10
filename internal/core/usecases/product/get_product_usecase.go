package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/dto"
	"tech-challenge-2-app-example/internal/core/port"
)

type getProductUseCase struct {
	repository port.ProductRepository
	presenter  port.ProductPresenter
}

func NewGetProductUseCase(repo port.ProductRepository, presenter port.ProductPresenter) port.GetProductUseCase {
	return &getProductUseCase{
		repository: repo,
		presenter:  presenter,
	}
}

func (uc *getProductUseCase) Execute(ctx context.Context, id uint64) (*dto.ProductResponse, error) {
	product, err := uc.repository.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if product == nil {
		return nil, errors.NewNotFoundError("Produto não encontrado")
	}

	response := uc.presenter.ToResponse(product)
	return &response, nil
}
