package repository

import (
	"context"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/port"
)

type productRepository struct {
	dataSource port.ProductDataSource
}

func NewProductRepository(dataSource port.ProductDataSource) port.ProductRepository {
	return &productRepository{
		dataSource: dataSource,
	}
}

func (r *productRepository) FindByID(ctx context.Context, id uint64) (*entity.Product, error) {
	return r.dataSource.FindByID(ctx, id)
}

func (r *productRepository) FindAll(ctx context.Context, name string, categoryID uint64, page, limit int) ([]*entity.Product, int64, error) {
	filters := make(map[string]interface{})

	if name != "" {
		filters["name"] = name
	}
	if categoryID != 0 {
		filters["category_id"] = categoryID
	}

	return r.dataSource.FindAll(ctx, filters, page, limit)
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) error {
	return r.dataSource.Create(ctx, product)
}

func (r *productRepository) Update(ctx context.Context, product *entity.Product) error {
	return r.dataSource.Update(ctx, product)
}

func (r *productRepository) Delete(ctx context.Context, id uint64) error {
	return r.dataSource.Delete(ctx, id)
}
