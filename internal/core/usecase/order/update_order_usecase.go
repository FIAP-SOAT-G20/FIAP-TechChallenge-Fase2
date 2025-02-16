package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateOrderUseCase struct {
	gateway   port.OrderGateway
	presenter port.OrderPresenter
}

// NewUpdateOrderUseCase creates a new UpdateOrderUseCase
func NewUpdateOrderUseCase(gateway port.OrderGateway, presenter port.OrderPresenter) port.UpdateOrderUseCase {
	return &updateOrderUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute updates a order
func (uc *updateOrderUseCase) Execute(ctx context.Context, input dto.UpdateOrderInput) error {
	order, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if order == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	if input.CustomerID != 0 && order.CustomerID != input.CustomerID {
		return domain.NewInvalidInputError(domain.ErrInvalidBody)
	}

	if input.Status != "" && order.Status != input.Status {
		if !entity.CanTransitionTo(order.Status, input.Status) {
			return domain.NewInvalidInputError(domain.ErrInvalidBody)
		}

		// TODO: Implement this validation when staff is implemented
		// if entity.StatusTransitionNeedsStaffID(order.Status) && staffID == nil {
		// 	return domain.NewInternalError(errors.New(domain.ErrOrderMandatoryStaffId))
		// }

		// if order.Status == entity.PENDING && len(order.OrderProducts) == 0 {
		// 	return domain.ErrOrderWithoutProducts
		// }
	}

	order.Update(input.CustomerID, input.Status)

	if err := uc.gateway.Update(ctx, order); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})
	return nil
}
