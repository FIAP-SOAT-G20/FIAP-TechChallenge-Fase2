package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type OrderGateway interface {
	FindByID(ctx context.Context, id uint64) (*entity.Order, error)
	FindAll(ctx context.Context, customerId uint64, status string, page, limit int) ([]*entity.Order, int64, error)
	Create(ctx context.Context, order *entity.Order) error
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, id uint64) error
}
