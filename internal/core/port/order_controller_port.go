package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type OrderController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListOrdersInput) error
	Create(ctx context.Context, presenter Presenter, input dto.CreateOrderInput) error
	Get(ctx context.Context, presenter Presenter, input dto.GetOrderInput) error
	Update(ctx context.Context, presenter Presenter, input dto.UpdateOrderInput) error
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteOrderInput) error
}
