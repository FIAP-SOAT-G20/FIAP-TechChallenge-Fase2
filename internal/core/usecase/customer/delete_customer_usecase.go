package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteCustomerUseCase struct {
	gateway   port.CustomerGateway
	presenter port.CustomerPresenter
}

// NewDeleteCustomerUseCase creates a new DeleteCustomerUseCase
func NewDeleteCustomerUseCase(gateway port.CustomerGateway, presenter port.CustomerPresenter) port.DeleteCustomerUseCase {
	return &deleteCustomerUseCase{gateway, presenter}
}

// Execute deletes a customer
func (uc *deleteCustomerUseCase) Execute(ctx context.Context, input dto.DeleteCustomerInput) error {
	customer, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if customer == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})

	return nil
}
