package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteOrderProductUseCase struct {
	gateway port.OrderProductGateway
}

// NewDeleteOrderProductUseCase creates a new DeleteOrderProductUseCase
func NewDeleteOrderProductUseCase(gateway port.OrderProductGateway) port.DeleteOrderProductUseCase {
	return &deleteOrderProductUseCase{gateway}
}

// Execute deletes a order
func (uc *deleteOrderProductUseCase) Execute(ctx context.Context, input dto.DeleteOrderProductInput) (*entity.OrderProduct, error) {
	order, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.OrderID, input.ProductID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}
