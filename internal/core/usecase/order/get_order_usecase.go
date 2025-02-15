package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getOrderUseCase struct {
	gateway   port.OrderGateway
	presenter port.OrderPresenter
}

// NewGetOrderUseCase creates a new GetOrderUseCase
func NewGetOrderUseCase(gateway port.OrderGateway, presenter port.OrderPresenter) port.GetOrderUseCase {
	return &getOrderUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute gets a order
func (uc *getOrderUseCase) Execute(ctx context.Context, input dto.GetOrderInput) error {
	order, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if order == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	uc.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})
	return nil
}
