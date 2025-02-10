package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
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
		return domain.NewInternalError(err)
	}
	if product == nil {
		return domain.NewNotFoundError("produto não encontrado")
	}

	// Deleta o produto
	if err := uc.gateway.Delete(ctx, id); err != nil {
		return domain.NewInternalError(err)
	}

	return nil
}
