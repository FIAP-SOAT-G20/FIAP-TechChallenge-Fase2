package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateOrderUseCase struct {
	gateway port.OrderGateway
}

// NewUpdateOrderUseCase creates a new UpdateOrderUseCase
func NewUpdateOrderUseCase(gateway port.OrderGateway) port.UpdateOrderUseCase {
	return &updateOrderUseCase{gateway}
}

// Execute updates a order
func (uc *updateOrderUseCase) Execute(ctx context.Context, input dto.UpdateOrderInput) (*entity.Order, error) {
	order, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if input.CustomerID != 0 && order.CustomerID != input.CustomerID {
		return nil, domain.NewInvalidInputError(domain.ErrInvalidBody)
	}

	if input.Status != "" && order.Status != input.Status {
		if !valueobject.StatusCanTransitionTo(order.Status, input.Status) {
			return nil, domain.NewInvalidInputError(domain.ErrInvalidBody)
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
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}
