package usecase

import (
	"context"
	"errors"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderHistoryUseCase struct {
	gateway port.OrderHistoryGateway
}

// NewOrderHistoryUseCase creates a new OrderHistoryUseCase
func NewOrderHistoryUseCase(gateway port.OrderHistoryGateway) port.OrderHistoryUseCase {
	return &orderHistoryUseCase{gateway: gateway}
}

// List returns a list of orderHistories
func (uc *orderHistoryUseCase) List(ctx context.Context, input dto.ListOrderHistoriesInput) ([]*entity.OrderHistory, int64, error) {
	orderHistories, total, err := uc.gateway.FindAll(ctx, input.OrderID, input.Status, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return orderHistories, total, nil
}

// Create creates a new orderHistory
func (uc *orderHistoryUseCase) Create(ctx context.Context, input dto.CreateOrderHistoryInput) (*entity.OrderHistory, error) {
	if input.OrderID == 0 {
		return nil, domain.NewValidationError(errors.New("Invalid order id "))
	}

	orderStatus := valueobject.ToOrderStatus(input.Status)
	if orderStatus == valueobject.UNDEFINDED {
		return nil, domain.NewValidationError(errors.New("Invalid order status"))
	}

	orderHistory := entity.NewOrderHistory(input.OrderID, orderStatus, input.StaffID)

	if err := uc.gateway.Create(ctx, orderHistory); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return orderHistory, nil
}

// Get returns a orderHistory by ID
func (uc *orderHistoryUseCase) Get(ctx context.Context, input dto.GetOrderHistoryInput) (*entity.OrderHistory, error) {
	orderHistory, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if orderHistory == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return orderHistory, nil
}

// Delete deletes a orderHistory
func (uc *orderHistoryUseCase) Delete(ctx context.Context, input dto.DeleteOrderHistoryInput) (*entity.OrderHistory, error) {
	orderHistory, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if orderHistory == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return orderHistory, nil
}
