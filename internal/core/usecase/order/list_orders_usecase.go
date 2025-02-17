package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listOrdersUseCase struct {
	gateway   port.OrderGateway
	presenter port.OrderPresenter
}

// NewListOrdersUseCase creates a new ListOrdersUseCase
func NewListOrdersUseCase(gateway port.OrderGateway, presenter port.OrderPresenter) port.ListOrdersUseCase {
	return &listOrdersUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute lists all orders
func (uc *listOrdersUseCase) Execute(ctx context.Context, input dto.ListOrdersInput) error {
	orders, total, err := uc.gateway.FindAll(ctx, input.CustomerID, input.Status, input.Page, input.Limit)
	if err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orders,
	})
	return nil
}
