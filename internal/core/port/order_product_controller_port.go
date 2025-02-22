package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type OrderProductController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListOrderProductsInput) error
	Create(ctx context.Context, presenter Presenter, input dto.CreateOrderProductInput) error
	Get(ctx context.Context, presenter Presenter, input dto.GetOrderProductInput) error
	Update(ctx context.Context, presenter Presenter, input dto.UpdateOrderProductInput) error
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteOrderProductInput) error
}
