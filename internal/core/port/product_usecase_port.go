package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type ProductUseCase interface {
	List(ctx context.Context, input dto.ListProductsInput) ([]*entity.Product, int64, error)
	Create(ctx context.Context, input dto.CreateProductInput) (*entity.Product, error)
	Get(ctx context.Context, input dto.GetProductInput) (*entity.Product, error)
	Update(ctx context.Context, input dto.UpdateProductInput) (*entity.Product, error)
	Delete(ctx context.Context, input dto.DeleteProductInput) (*entity.Product, error)
}
