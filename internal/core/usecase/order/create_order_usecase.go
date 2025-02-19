package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createOrderUseCase struct {
	gateway port.OrderGateway
}

// NewCreateOrderUseCase creates a new CreateOrderUseCase
func NewCreateOrderUseCase(gateway port.OrderGateway) port.CreateOrderUseCase {
	return &createOrderUseCase{gateway}
}

// Execute creates a new Order
func (uc *createOrderUseCase) Execute(ctx context.Context, input dto.CreateOrderInput) (*entity.Order, error) {
	order := entity.NewOrder(input.CustomerID)

	if err := uc.gateway.Create(ctx, order); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}
