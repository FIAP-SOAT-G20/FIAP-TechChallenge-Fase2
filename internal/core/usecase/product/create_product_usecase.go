package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/port"
	"tech-challenge-2-app-example/internal/core/usecase"
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

func (uc *createProductUseCase) Execute(ctx context.Context, input usecase.CreateProductInput) (*usecase.ProductOutput, error) {
	// Cria e valida o produto usando as regras de domínio
	product, err := entity.NewProduct(input.Name, input.Description, input.Price, input.CategoryID)
	if err != nil {
		return nil, errors.NewValidationError(err)
	}

	// Persiste o produto
	if err := uc.gateway.Create(ctx, product); err != nil {
		return nil, errors.NewInternalError(err)
	}

	// Formata a resposta usando o presenter
	output := uc.presenter.ToOutput(product)
	return output, nil
}
