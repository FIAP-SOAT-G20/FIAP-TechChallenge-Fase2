package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createCustomerUseCase struct {
	gateway   port.CustomerGateway
	presenter port.CustomerPresenter
}

// NewCreateCustomerUseCase creates a new CreateCustomerUseCase
func NewCreateCustomerUseCase(gateway port.CustomerGateway, presenter port.CustomerPresenter) port.CreateCustomerUseCase {
	return &createCustomerUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute creates a new Customer
func (uc *createCustomerUseCase) Execute(ctx context.Context, input dto.CreateCustomerInput) error {
	customer := entity.NewCustomer(input.Name, input.Email, input.CPF)

	if err := uc.gateway.Create(ctx, customer); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})
	return nil
}
