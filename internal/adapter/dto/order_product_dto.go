package dto

import "time"

type CreateOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
	Quantity  uint32
	// Order     Order
	// Product   Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
	Quantity  uint32
}

type GetOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
}

type DeleteOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
}

type ListOrderProductsInput struct {
	OrderID   uint64
	ProductID uint64
	Page      int
	Limit     int
}

type OrderProductPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
}
