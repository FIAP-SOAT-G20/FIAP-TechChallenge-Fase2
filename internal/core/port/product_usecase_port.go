package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type ListProductsUseCase interface {
	Execute(ctx context.Context, input dto.ListProductsInput) ([]*entity.Product, int64, error)
}

type CreateProductUseCase interface {
	Execute(ctx context.Context, input dto.CreateProductInput) (*entity.Product, error)
}

type GetProductUseCase interface {
	Execute(ctx context.Context, input dto.GetProductInput) (*entity.Product, error)
}

type UpdateProductUseCase interface {
	Execute(ctx context.Context, input dto.UpdateProductInput) (*entity.Product, error)
}

type DeleteProductUseCase interface {
	Execute(ctx context.Context, input dto.DeleteProductInput) (*entity.Product, error)
}
