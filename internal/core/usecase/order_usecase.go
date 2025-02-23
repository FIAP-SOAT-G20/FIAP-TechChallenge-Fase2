package usecase

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
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
func (uc *orderUseCase) List(ctx context.Context, i dto.ListOrdersInput) ([]*entity.Order, int64, error) {
	orders, total, err := uc.gateway.FindAll(ctx, i.CustomerID, i.Status, i.Page, i.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return orders, total, nil
}

// Create creates a new Order
func (uc *orderUseCase) Create(ctx context.Context, i dto.CreateOrderInput) (*entity.Order, error) {
	order := &entity.Order{CustomerID: i.CustomerID}

	if err := uc.gateway.Create(ctx, order); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}

// Get returns a Order by ID
func (uc *orderUseCase) Get(ctx context.Context, i dto.GetOrderInput) (*entity.Order, error) {
	order, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return order, nil
}

// Update updates a Order
func (uc *orderUseCase) Update(ctx context.Context, i dto.UpdateOrderInput) (*entity.Order, error) {
	order, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if i.CustomerID != 0 && order.CustomerID != i.CustomerID {
		return nil, domain.NewInvalidInputError(domain.ErrInvalidBody)
	}

	if i.Status != "" && order.Status != i.Status {
		if !valueobject.StatusCanTransitionTo(order.Status, i.Status) {
			return nil, domain.NewInvalidInputError(domain.ErrInvalidBody)
		}

		if valueobject.StatusTransitionNeedsStaffID(i.Status) && i.StaffID == 0 {
			return nil, domain.NewInvalidInputError(domain.ErrStaffIdIsMandatory)
		}
	}

	orderProducts := order.OrderProducts
	order.Update(i.CustomerID, i.Status)

	if err := uc.gateway.Update(ctx, order); err != nil {
		return nil, domain.NewInternalError(err)
	}

	// Restore order products, to calculate total bill in the presenter
	order.OrderProducts = orderProducts

	return order, nil
}

// Delete deletes a Order
func (uc *orderUseCase) Delete(ctx context.Context, i dto.DeleteOrderInput) (*entity.Order, error) {
	order, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if order == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, i.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return order, nil
}
