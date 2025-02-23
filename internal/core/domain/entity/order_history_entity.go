package entity

import (
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"time"
)

type OrderHistory struct {
	ID        uint64
	OrderID   uint64
	StaffID   *uint64
	Status    valueobject.OrderStatus
	CreatedAt time.Time
	Order     Order
	Staff     *Staff
}

func NewOrderHistory(orderID uint64, status valueobject.OrderStatus, staffID *uint64) *OrderHistory {
	return &OrderHistory{
		OrderID: orderID,
		Status:  status,
		StaffID: staffID,
	}
}
