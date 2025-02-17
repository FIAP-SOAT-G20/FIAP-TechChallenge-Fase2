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

func (pg *orderGateway) FindByID(ctx context.Context, id uint64) (*entity.Order, error) {
	return pg.dataSource.FindByID(ctx, id)
}

func (pg *orderGateway) FindAll(ctx context.Context, customerId uint64, status valueobject.OrderStatus, page, limit int) ([]*entity.Order, int64, error) {
	filters := make(map[string]interface{})

	if customerId != 0 {
		filters["customer_id"] = customerId
	}

	if status != "" {
		filters["status"] = status
	}

	return pg.dataSource.FindAll(ctx, filters, page, limit)
}

func (pg *orderGateway) Create(ctx context.Context, order *entity.Order) error {
	return pg.dataSource.Create(ctx, order)
}

func (pg *orderGateway) Update(ctx context.Context, order *entity.Order) error {
	return pg.dataSource.Update(ctx, order)
}

func (pg *orderGateway) Delete(ctx context.Context, id uint64) error {
	return pg.dataSource.Delete(ctx, id)
}
