package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderGateway struct {
	dataSource port.OrderDataSource
}

func NewOrderGateway(dataSource port.OrderDataSource) port.OrderGateway {
	return &orderGateway{
		dataSource: dataSource,
	}
}

func (g *orderGateway) FindByID(ctx context.Context, id uint64) (*entity.Order, error) {
	return g.dataSource.FindByID(ctx, id)
}

func (g *orderGateway) FindAll(ctx context.Context, customerId uint64, status valueobject.OrderStatus, page, limit int) ([]*entity.Order, int64, error) {
	filters := make(map[string]interface{})

	if customerId != 0 {
		filters["customer_id"] = customerId
	}

	if status != "" {
		filters["status"] = status
	}

	return g.dataSource.FindAll(ctx, filters, page, limit)
}

func (g *orderGateway) Create(ctx context.Context, order *entity.Order) error {
	return g.dataSource.Create(ctx, order)
}

func (g *orderGateway) Update(ctx context.Context, order *entity.Order) error {
	return g.dataSource.Update(ctx, order)
}

func (g *orderGateway) Delete(ctx context.Context, id uint64) error {
	return g.dataSource.Delete(ctx, id)
}
