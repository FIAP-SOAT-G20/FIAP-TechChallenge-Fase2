package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type ListOrdersUseCase interface {
	Execute(ctx context.Context, input dto.ListOrdersInput) ([]*entity.Order, int64, error)
}

type CreateOrderUseCase interface {
	Execute(ctx context.Context, input dto.CreateOrderInput) (*entity.Order, error)
}

type GetOrderUseCase interface {
	Execute(ctx context.Context, input dto.GetOrderInput) (*entity.Order, error)
}

type UpdateOrderUseCase interface {
	Execute(ctx context.Context, input dto.UpdateOrderInput) (*entity.Order, error)
}

type DeleteOrderUseCase interface {
	Execute(ctx context.Context, input dto.DeleteOrderInput) (*entity.Order, error)
}
