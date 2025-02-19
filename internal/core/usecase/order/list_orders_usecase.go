package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listOrdersUseCase struct {
	gateway port.OrderGateway
}

// NewListOrdersUseCase creates a new ListOrdersUseCase
func NewListOrdersUseCase(gateway port.OrderGateway) port.ListOrdersUseCase {
	return &listOrdersUseCase{gateway}
}

// Execute lists all orders
func (uc *listOrdersUseCase) Execute(ctx context.Context, input dto.ListOrdersInput) ([]*entity.Order, int64, error) {
	orders, total, err := uc.gateway.FindAll(ctx, input.CustomerID, input.Status, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return orders, total, nil
}
