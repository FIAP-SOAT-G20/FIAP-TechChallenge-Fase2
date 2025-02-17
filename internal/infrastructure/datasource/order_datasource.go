package datasource

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderDataSource struct {
	db *gorm.DB
}

type orderKey string

func NewOrderDataSource(db *gorm.DB) port.OrderDataSource {
	return &orderDataSource{
		db: db,
	}
}

func (ds *orderDataSource) FindByID(ctx context.Context, id uint64) (*entity.Order, error) {
	var order entity.Order
	result := ds.db.WithContext(ctx).Preload("OrderProducts.Product").First(&order, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding order: %w", result.Error)
	}
	return &order, nil
}

func (ds *orderDataSource) FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]*entity.Order, int64, error) {
	var orders []*entity.Order
	var total int64

	query := ds.db.WithContext(ctx).Preload("OrderProducts.Product")

	// Apply filters
	for key, value := range filters {
		switch key {
		case "status":
			if status, ok := value.(entity.OrderStatus); ok && status != entity.UNDEFINDED {
				query = query.Where("status = ?", status)
			}
		case "customer_id":
			if customerID, ok := value.(uint64); ok && customerID != 0 {
				query = query.Where("customer_id = ?", customerID)
			}
		}
	}

	// Count total before pagination
	if err := query.Model(&entity.Order{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting orders: %w", err)
	}

	// Get paginated results
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, fmt.Errorf("error finding orders: %w", err)
	}

	return orders, total, nil
}

func (ds *orderDataSource) Create(ctx context.Context, order *entity.Order) error {
	if err := ds.db.WithContext(ctx).Create(order).Error; err != nil {
		return fmt.Errorf("error creating order: %w", err)
	}
	return nil
}

func (ds *orderDataSource) Update(ctx context.Context, order *entity.Order) error {
	result := ds.db.WithContext(ctx).Model(order).Updates(order)
	if result.Error != nil {
		return fmt.Errorf("error updating order: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (ds *orderDataSource) Delete(ctx context.Context, id uint64) error {
	result := ds.db.WithContext(ctx).Delete(&entity.Order{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting order: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (ds *orderDataSource) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return ds.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create a new context with the transaction
		keyPrincipalID := orderKey(uuid.NewString())
		txCtx := context.WithValue(ctx, keyPrincipalID, tx)
		return fn(txCtx)
	})
}
