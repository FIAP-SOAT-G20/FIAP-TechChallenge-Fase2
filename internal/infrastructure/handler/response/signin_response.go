package response

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"

type SignInResponse struct {
	Customer CustomerResponse `json:"customer"`
}

type CustomerResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

func NewCustomerResponse(customer *entity.Customer) *CustomerResponse {
	if customer == nil {
		return nil
	}

	return &CustomerResponse{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
		CPF:   customer.CPF,
	}
}

func NewSignInResponse(customer *entity.Customer) *SignInResponse {
	if customer == nil {
		return nil
	}

	return &SignInResponse{
		Customer: *NewCustomerResponse(customer),
	}
}
