package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
)

type ListOrdersUseCase interface {
	Execute(ctx context.Context, input dto.ListOrdersInput) error
}

type CreateOrderUseCase interface {
	Execute(ctx context.Context, input dto.CreateOrderInput) error
}

type GetOrderUseCase interface {
	Execute(ctx context.Context, input dto.GetOrderInput) error
}

type UpdateOrderUseCase interface {
	Execute(ctx context.Context, input dto.UpdateOrderInput) error
}

type DeleteOrderUseCase interface {
	Execute(ctx context.Context, input dto.DeleteOrderInput) error
}
