package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteOrderUseCase struct {
	gateway port.OrderGateway
}

// NewDeleteOrderUseCase creates a new DeleteOrderUseCase
func NewDeleteOrderUseCase(gateway port.OrderGateway) port.DeleteOrderUseCase {
	return &deleteOrderUseCase{gateway}
}

// Execute deletes a order
func (uc *deleteOrderUseCase) Execute(ctx context.Context, input dto.DeleteOrderInput) (*entity.Order, error) {
	order, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}
