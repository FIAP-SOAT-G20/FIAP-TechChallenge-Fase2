package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
)

type ListOrderProductsUseCase interface {
	Execute(ctx context.Context, input dto.ListOrderProductsInput) error
}

type CreateOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.CreateOrderProductInput) error
}

type GetOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.GetOrderProductInput) error
}

type UpdateOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.UpdateOrderProductInput) error
}

type DeleteOrderProductUseCase interface {
	Execute(ctx context.Context, input dto.DeleteOrderProductInput) error
}
