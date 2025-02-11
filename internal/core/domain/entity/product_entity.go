package entity

import (
	"time"
)

type Product struct {
	ID          uint64
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(name string, description string, price float64, categoryID uint64) *Product {
	product := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return product
}

func (p *Product) Update(name string, description string, price float64, categoryID uint64) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.CategoryID = categoryID
	p.UpdatedAt = time.Now()
}
