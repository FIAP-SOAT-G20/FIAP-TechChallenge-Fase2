package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type ListOrderProductsUseCase interface {
	Execute(ctx context.Context, input dto.ListOrderProductsInput) ([]*entity.OrderProduct, int64, error)
}

type CreateOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.CreateOrderProductInput) (*entity.OrderProduct, error)
}

type GetOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.GetOrderProductInput) (*entity.OrderProduct, error)
}

type UpdateOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.UpdateOrderProductInput) (*entity.OrderProduct, error)
}

type DeleteOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.DeleteOrderProductInput) (*entity.OrderProduct, error)
}
