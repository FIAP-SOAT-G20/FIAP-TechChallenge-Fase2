package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createCustomerUseCase struct {
	gateway port.CustomerGateway
}

// NewCreateCustomerUseCase creates a new CreateCustomerUseCase
func NewCreateCustomerUseCase(gateway port.CustomerGateway) port.CreateCustomerUseCase {
	return &createCustomerUseCase{gateway}
}

// Execute creates a new Customer
func (uc *createCustomerUseCase) Execute(ctx context.Context, input dto.CreateCustomerInput) (*entity.Customer, error) {
	customer := entity.NewCustomer(input.Name, input.Email, input.CPF)

	if err := uc.gateway.Create(ctx, customer); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return customer, nil
}
