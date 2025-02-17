package dto

import (
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
)

type CreateOrderInput struct {
	CustomerID uint64
	Writer     ResponseWriter
}

type UpdateOrderInput struct {
	ID         uint64
	CustomerID uint64
	Status     valueobject.OrderStatus
	// Payment       Payment
	Writer ResponseWriter
}

type GetOrderInput struct {
	ID     uint64
	Writer ResponseWriter
}

type DeleteOrderInput struct {
	ID     uint64
	Writer ResponseWriter
}

type ListOrdersInput struct {
	CustomerID uint64
	Status     valueobject.OrderStatus
	Page       int
	Limit      int
	Writer     ResponseWriter
}

type OrderPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
	Writer ResponseWriter
}
