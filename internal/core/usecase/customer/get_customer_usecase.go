package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getCustomerUseCase struct {
	gateway   port.CustomerGateway
	presenter port.CustomerPresenter
}

// NewGetCustomerUseCase creates a new GetCustomerUseCase
func NewGetCustomerUseCase(gateway port.CustomerGateway, presenter port.CustomerPresenter) port.GetCustomerUseCase {
	return &getCustomerUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute gets a customer
func (uc *getCustomerUseCase) Execute(ctx context.Context, input dto.GetCustomerInput) error {
	customer, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if customer == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	uc.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})
	return nil
}
