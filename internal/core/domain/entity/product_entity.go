package entity

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/validator"
)

type Product struct {
	ID          uint64
	Name        string  `validate:"required,min=3,max=100"`
	Description string  `validate:"max=500"`
	Price       float64 `validate:"required,gt=0"`
	CategoryID  uint64  `validate:"required,gt=0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(name string, description string, price float64, categoryID uint64) (*Product, error) {
	product := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Update(name string, description string, price float64, categoryID uint64) error {
	p.Name = name
	p.Description = description
	p.Price = price
	p.CategoryID = categoryID
	p.UpdatedAt = time.Now()

	return p.Validate()
}

func (p *Product) Validate() error {
	return validator.GetValidator().Struct(p)
}
