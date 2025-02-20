package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderUseCase struct {
	gateway port.OrderGateway
}

// NewOrderUseCase creates a new OrdersUseCase
func NewOrderUseCase(gateway port.OrderGateway) port.OrderUseCase {
	return &orderUseCase{gateway}
}

// List returns a list of Orders
func (uc *orderUseCase) List(ctx context.Context, input dto.ListOrdersInput) ([]*entity.Order, int64, error) {
	orders, total, err := uc.gateway.FindAll(ctx, input.CustomerID, input.Status, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return orders, total, nil
}

// Create creates a new Order
func (uc *orderUseCase) Create(ctx context.Context, input dto.CreateOrderInput) (*entity.Order, error) {
	order := entity.NewOrder(input.CustomerID)

	if err := uc.gateway.Create(ctx, order); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}

// Get returns a Order by ID
func (uc *orderUseCase) Get(ctx context.Context, input dto.GetOrderInput) (*entity.Order, error) {
	order, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return order, nil
}

// Update updates a Order
func (uc *orderUseCase) Update(ctx context.Context, input dto.UpdateOrderInput) (*entity.Order, error) {
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

// Delete deletes a Order
func (uc *orderUseCase) Delete(ctx context.Context, input dto.DeleteOrderInput) (*entity.Order, error) {
	order, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}
