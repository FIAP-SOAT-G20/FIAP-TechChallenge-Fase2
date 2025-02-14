package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateCustomerUseCase struct {
	gateway   port.CustomerGateway
	presenter port.CustomerPresenter
}

// NewUpdateCustomerUseCase creates a new UpdateCustomerUseCase
func NewUpdateCustomerUseCase(gateway port.CustomerGateway, presenter port.CustomerPresenter) port.UpdateCustomerUseCase {
	return &updateCustomerUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute updates a customer
func (uc *updateCustomerUseCase) Execute(ctx context.Context, input dto.UpdateCustomerInput) error {
	customer, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if customer == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	customer.Update(input.Name, input.Email)

	if err := uc.gateway.Update(ctx, customer); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})
	return nil
}
