package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type OrderDataSource interface {
	FindByID(ctx context.Context, id uint64) (*entity.Order, error)
	FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]*entity.Order, int64, error)
	Create(ctx context.Context, order *entity.Order) error
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, id uint64) error
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}
