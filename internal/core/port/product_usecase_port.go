package port

import (
	"context"

	"tech-challenge-2-app-example/internal/core/usecase"
)

type ListProductsUseCase interface {
	Execute(ctx context.Context, input usecase.ListProductsInput) (*usecase.ListProductPaginatedOutput, error)
}

type CreateProductUseCase interface {
	Execute(ctx context.Context, input usecase.CreateProductInput) (*usecase.ProductOutput, error)
}

type GetProductUseCase interface {
	Execute(ctx context.Context, id uint64) (*usecase.ProductOutput, error)
}

type UpdateProductUseCase interface {
	Execute(ctx context.Context, id uint64, input usecase.UpdateProductInput) (*usecase.ProductOutput, error)
}

type DeleteProductUseCase interface {
	Execute(ctx context.Context, id uint64) error
}
