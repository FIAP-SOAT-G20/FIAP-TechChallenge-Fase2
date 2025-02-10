package port

import (
	"context"

	"tech-challenge-2-app-example/internal/core/dto"
)

type ListProductsUseCase interface {
	Execute(ctx context.Context, req dto.ProductListRequest) (*dto.PaginatedResponse, error)
}

type CreateProductUseCase interface {
	Execute(ctx context.Context, req dto.ProductRequest) (*dto.ProductResponse, error)
}

type GetProductUseCase interface {
	Execute(ctx context.Context, id uint64) (*dto.ProductResponse, error)
}

type UpdateProductUseCase interface {
	Execute(ctx context.Context, id uint64, req dto.ProductRequest) (*dto.ProductResponse, error)
}

type DeleteProductUseCase interface {
	Execute(ctx context.Context, id uint64) error
}
