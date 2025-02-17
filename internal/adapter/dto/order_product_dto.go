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
	Writer    ResponseWriter
}

type UpdateOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
	Quantity  uint32
	Writer    ResponseWriter
}

type GetOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
	Writer    ResponseWriter
}

type DeleteOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
	Writer    ResponseWriter
}

type ListOrderProductsInput struct {
	OrderID   uint64
	ProductID uint64
	Page      int
	Limit     int
	Writer    ResponseWriter
}

type OrderProductPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
	Writer ResponseWriter
}
