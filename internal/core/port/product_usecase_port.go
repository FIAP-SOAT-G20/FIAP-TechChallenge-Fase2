package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
)

type ListProductsUseCase interface {
	Execute(ctx context.Context, input dto.ListProductsInput) error
}

type CreateProductUseCase interface {
	Execute(ctx context.Context, input dto.CreateProductInput) error
}

type GetProductUseCase interface {
	Execute(ctx context.Context, input dto.GetProductInput) error
}

type UpdateProductUseCase interface {
	Execute(ctx context.Context, input dto.UpdateProductInput) error
}

type DeleteProductUseCase interface {
	Execute(ctx context.Context, input dto.DeleteProductInput) error
}
