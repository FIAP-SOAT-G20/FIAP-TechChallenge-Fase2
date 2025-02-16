package entity

import (
	"time"
)

type OrderProduct struct {
	OrderID   uint64
	ProductID uint64
	Quantity  uint32
	// Order     Order // Mover para response
	// Product   Product // Mover para response
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrderProduct(orderID uint64, productID uint64, quantity uint32) *OrderProduct {
	orderProduct := &OrderProduct{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return orderProduct
}

func (p *OrderProduct) Update(quantity uint32) {
	p.Quantity = quantity
	p.UpdatedAt = time.Now()
}
