package product

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/port"
)

type deleteProductUseCase struct {
	gateway port.ProductGateway
}

func NewDeleteProductUseCase(gateway port.ProductGateway) port.DeleteProductUseCase {
	return &deleteProductUseCase{
		gateway: gateway,
	}
}

func (uc *deleteProductUseCase) Execute(ctx context.Context, id uint64) error {
	// Verifica se o produto existe
	product, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return errors.NewInternalError(err)
	}
	if product == nil {
		return errors.NewNotFoundError("produto não encontrado")
	}

	// Deleta o produto
	if err := uc.gateway.Delete(ctx, id); err != nil {
		return errors.NewInternalError(err)
	}

	return nil
}
