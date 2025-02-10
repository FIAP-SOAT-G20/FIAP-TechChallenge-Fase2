package gateway

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/port"
)

type productGateway struct {
	dataSource port.ProductDataSource
}

func NewProductGateway(dataSource port.ProductDataSource) port.ProductGateway {
	return &productGateway{
		dataSource: dataSource,
	}
}

func (r *productGateway) FindByID(ctx context.Context, id uint64) (*entity.Product, error) {
	return r.dataSource.FindByID(ctx, id)
}

func (r *productGateway) FindAll(ctx context.Context, name string, categoryID uint64, page, limit int) ([]*entity.Product, int64, error) {
	filters := make(map[string]interface{})

	if name != "" {
		filters["name"] = name
	}
	if categoryID != 0 {
		filters["category_id"] = categoryID
	}

	return r.dataSource.FindAll(ctx, filters, page, limit)
}

func (r *productGateway) Create(ctx context.Context, product *entity.Product) error {
	return r.dataSource.Create(ctx, product)
}

func (r *productGateway) Update(ctx context.Context, product *entity.Product) error {
	return r.dataSource.Update(ctx, product)
}

func (r *productGateway) Delete(ctx context.Context, id uint64) error {
	return r.dataSource.Delete(ctx, id)
}
