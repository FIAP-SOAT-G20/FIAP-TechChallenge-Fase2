package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/dto"
	"tech-challenge-2-app-example/internal/core/port"
)

type createProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

func NewCreateProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.CreateProductUseCase {
	return &createProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

func (uc *createProductUseCase) Execute(ctx context.Context, req dto.ProductRequest) (*dto.ProductResponse, error) {
	// Cria e valida o produto usando as regras de domínio
	product, err := entity.NewProduct(req.Name, req.Description, req.Price, req.CategoryID)
	if err != nil {
		return nil, errors.NewValidationError(err)
	}

	// Persiste o produto
	if err := uc.gateway.Create(ctx, product); err != nil {
		return nil, errors.NewInternalError(err)
	}

	// Formata a resposta
	response := uc.presenter.ToResponse(product)
	return &response, nil
}
