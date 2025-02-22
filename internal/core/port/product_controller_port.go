package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type ProductController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListProductsInput) error
	Create(ctx context.Context, presenter Presenter, input dto.CreateProductInput) error
	Get(ctx context.Context, presenter Presenter, input dto.GetProductInput) error
	Update(ctx context.Context, presenter Presenter, input dto.UpdateProductInput) error
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteProductInput) error
}
