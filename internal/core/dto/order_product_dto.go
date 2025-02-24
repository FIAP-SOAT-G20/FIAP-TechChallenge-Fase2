package dto

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type CreateOrderProductInput struct {
	OrderID   uint64
	ProductID uint64
	Quantity  uint32
}

func (i CreateOrderProductInput) ToEntity() *entity.OrderProduct {
	return &entity.OrderProduct{
		OrderID:   i.OrderID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
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
