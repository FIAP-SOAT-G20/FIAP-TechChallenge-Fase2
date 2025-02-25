package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type OrderHistoryController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListOrderHistoriesInput) error
	Create(ctx context.Context, presenter Presenter, input dto.CreateOrderHistoryInput) error
	Get(ctx context.Context, presenter Presenter, input dto.GetOrderHistoryInput) error
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteOrderHistoryInput) error
}
