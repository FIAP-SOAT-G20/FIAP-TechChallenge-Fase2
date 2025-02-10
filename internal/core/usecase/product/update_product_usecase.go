package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/port"
	"tech-challenge-2-app-example/internal/core/usecase"
)

type updateProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

func NewUpdateProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.UpdateProductUseCase {
	return &updateProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

func (uc *updateProductUseCase) Execute(ctx context.Context, id uint64, input usecase.UpdateProductInput) (*usecase.ProductOutput, error) {
	// Busca o produto existente
	product, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if product == nil {
		return nil, errors.NewNotFoundError("produto não encontrado")
	}

	// Atualiza o produto
	if err := product.Update(input.Name, input.Description, input.Price, input.CategoryID); err != nil {
		return nil, errors.NewValidationError(err)
	}

	// Persiste as alterações
	if err := uc.gateway.Update(ctx, product); err != nil {
		return nil, errors.NewInternalError(err)
	}

	output := uc.presenter.ToOutput(product)
	return output, nil
}
