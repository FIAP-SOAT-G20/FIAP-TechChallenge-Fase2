package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listCustomersUseCase struct {
	gateway   port.CustomerGateway
	presenter port.CustomerPresenter
}

// NewListCustomersUseCase creates a new ListCustomersUseCase
func NewListCustomersUseCase(gateway port.CustomerGateway, presenter port.CustomerPresenter) port.ListCustomersUseCase {
	return &listCustomersUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute lists all customers
func (uc *listCustomersUseCase) Execute(ctx context.Context, input dto.ListCustomersInput) error {
	customers, total, err := uc.gateway.FindAll(ctx, input.Name, input.Page, input.Limit)
	if err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: customers,
	})
	return nil
}
