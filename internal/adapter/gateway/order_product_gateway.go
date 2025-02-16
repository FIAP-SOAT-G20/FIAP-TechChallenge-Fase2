package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderProductGateway struct {
	dataSource port.OrderProductDataSource
}

func NewOrderProductGateway(dataSource port.OrderProductDataSource) port.OrderProductGateway {
	return &orderProductGateway{
		dataSource: dataSource,
	}
}

func (pg *orderProductGateway) FindByID(ctx context.Context, orderId, productId uint64) (*entity.OrderProduct, error) {
	return pg.dataSource.FindByID(ctx, orderId, productId)
}

func (pg *orderProductGateway) FindAll(ctx context.Context, orderId uint64, productId uint64, page, limit int) ([]*entity.OrderProduct, int64, error) {
	filters := make(map[string]interface{})

	if orderId != 0 {
		filters["order_id"] = orderId
	}

	if productId != 0 {
		filters["product_id"] = productId
	}

	return pg.dataSource.FindAll(ctx, filters, page, limit)
}

func (pg *orderProductGateway) Create(ctx context.Context, orderProduct *entity.OrderProduct) error {
	return pg.dataSource.Create(ctx, orderProduct)
}

func (pg *orderProductGateway) Update(ctx context.Context, orderProduct *entity.OrderProduct) error {
	return pg.dataSource.Update(ctx, orderProduct)
}

func (pg *orderProductGateway) Delete(ctx context.Context, orderId, productId uint64) error {
	return pg.dataSource.Delete(ctx, orderId, productId)
}
