package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type OrderProductUseCase interface {
	List(ctx context.Context, input dto.ListOrderProductsInput) ([]*entity.OrderProduct, int64, error)
	Create(ctx context.Context, input dto.CreateOrderProductInput) (*entity.OrderProduct, error)
	Get(ctx context.Context, input dto.GetOrderProductInput) (*entity.OrderProduct, error)
	Update(ctx context.Context, input dto.UpdateOrderProductInput) (*entity.OrderProduct, error)
	Delete(ctx context.Context, input dto.DeleteOrderProductInput) (*entity.OrderProduct, error)
}
