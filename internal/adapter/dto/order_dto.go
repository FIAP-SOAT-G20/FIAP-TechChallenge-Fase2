package dto

import (
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
)

type CreateOrderInput struct {
	CustomerID uint64
}

type UpdateOrderInput struct {
	ID         uint64
	CustomerID uint64
	Status     valueobject.OrderStatus
	// Payment       Payment
}

type GetOrderInput struct {
	ID uint64
}

type DeleteOrderInput struct {
	ID uint64
}

type ListOrdersInput struct {
	CustomerID uint64
	Status     valueobject.OrderStatus
	Page       int
	Limit      int
}
