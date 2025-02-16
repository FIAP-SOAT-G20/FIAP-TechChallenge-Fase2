package entity

import (
	"time"
)

type OrderProduct struct {
	OrderID   uint64
	ProductID uint64
	Price     float32
	Quantity  uint32
	// Order     Order // Mover para response
	// Product   Product // Mover para response
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrderProduct(orderID uint64, productID uint64, price float32, quantity uint32) *OrderProduct {
	orderProduct := &OrderProduct{
		OrderID:   orderID,
		ProductID: productID,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return orderProduct
}

func (p *OrderProduct) Update(price float32, quantity uint32) {
	p.Price = price
	p.Quantity = quantity
	p.UpdatedAt = time.Now()
}
