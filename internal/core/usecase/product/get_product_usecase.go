package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/port"
	"tech-challenge-2-app-example/internal/core/usecase"
)

type getProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

func NewGetProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.GetProductUseCase {
	return &getProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

func (uc *getProductUseCase) Execute(ctx context.Context, id uint64) (*usecase.ProductOutput, error) {
	product, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if product == nil {
		return nil, errors.NewNotFoundError("Produto não encontrado")
	}

	output := uc.presenter.ToOutput(product)
	return output, nil
}
