package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/dto"
	"tech-challenge-2-app-example/internal/core/port"
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

func (uc *getProductUseCase) Execute(ctx context.Context, id uint64) (*dto.ProductResponse, error) {
	product, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if product == nil {
		return nil, errors.NewNotFoundError("Produto não encontrado")
	}

	response := uc.presenter.ToResponse(product)
	return &response, nil
}
