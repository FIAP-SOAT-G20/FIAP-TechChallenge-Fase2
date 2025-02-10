package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/dto"
	"tech-challenge-2-app-example/internal/core/port"
)

type updateProductUseCase struct {
	repository port.ProductRepository
	presenter  port.ProductPresenter
}

func NewUpdateProductUseCase(repo port.ProductRepository, presenter port.ProductPresenter) port.UpdateProductUseCase {
	return &updateProductUseCase{
		repository: repo,
		presenter:  presenter,
	}
}

func (uc *updateProductUseCase) Execute(ctx context.Context, id uint64, req dto.ProductRequest) (*dto.ProductResponse, error) {
	// Busca o produto existente
	product, err := uc.repository.FindByID(ctx, id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if product == nil {
		return nil, errors.NewNotFoundError("produto não encontrado")
	}

	// Atualiza o produto
	if err := product.Update(req.Name, req.Description, req.Price, req.CategoryID); err != nil {
		return nil, errors.NewValidationError(err)
	}

	// Persiste as alterações
	if err := uc.repository.Update(ctx, product); err != nil {
		return nil, errors.NewInternalError(err)
	}

	response := uc.presenter.ToResponse(product)
	return &response, nil
}
