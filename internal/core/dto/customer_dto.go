package dto

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"

type CreateCustomerInput struct {
	Name  string
	Email string
	CPF   string
}

func (i CreateCustomerInput) ToEntity() *entity.Customer {
	return &entity.Customer{
		Name:  i.Name,
		Email: i.Email,
		CPF:   i.CPF,
	}
}

type UpdateCustomerInput struct {
	ID    uint64
	Name  string
	Email string
}

type GetCustomerInput struct {
	ID uint64
}

type DeleteCustomerInput struct {
	ID uint64
}

type ListCustomersInput struct {
	Name  string
	Page  int
	Limit int
}
