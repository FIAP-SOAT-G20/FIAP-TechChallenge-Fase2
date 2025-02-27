package dto

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
)

// Request DTOs
type SignInRequest struct {
	CPF string `json:"cpf" binding:"required" example:"000.000.000-00"`
}

// Response DTOs
type SignInResponse struct {
	CustomerResponse
}

func NewSignInResponse(customer *entity.Customer) *SignInResponse {
	if customer == nil {
		return nil
	}

	return &SignInResponse{
		CustomerResponse: *NewCustomerResponse(customer),
	}
} 