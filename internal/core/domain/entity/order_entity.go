package entity

import (
	"time"
)

type Order struct {
	ID         uint64
	CustomerID uint64
	TotalBill  float32
	Status     OrderStatus
	// Payment       Payment // Mover para response
	// Customer Customer // Mover para response
	OrderProducts []OrderProduct // Mover para response
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewOrder(consumerID uint64) *Order {
	order := &Order{
		CustomerID: consumerID,
		Status:     OPEN,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return order
}

func (p *Order) Update(customerID uint64, status OrderStatus) {
	if customerID != 0 {
		p.CustomerID = customerID
	}
	if status != UNDEFINDED {
		p.Status = status
	}
	p.OrderProducts = nil
	p.UpdatedAt = time.Now()
}
