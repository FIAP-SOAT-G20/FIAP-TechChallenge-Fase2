package entity

import (
	"time"
)

type Customer struct {
	ID        uint64
	Name      string
	Email     string
	CPF       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name string, email string, cpf string) *Customer {
	product := &Customer{
		Name:      name,
		Email:     email,
		CPF:       cpf,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return product
}

func (p *Customer) Update(name string, email string, cpf string) {
	p.Name = name
	p.Email = email
	p.CPF = cpf
	p.UpdatedAt = time.Now()
}
