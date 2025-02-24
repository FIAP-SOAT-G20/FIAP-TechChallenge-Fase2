package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
)

type paymentUseCase struct {
	orderGateway           port.OrderGateway
	paymentGateway         port.PaymentGateway
	paymentExternalGateway port.PaymentExternalGateway
}

// NewPaymentUseCase create a new payment use case
func NewPaymentUseCase(
	orderGateway port.OrderGateway,
	paymentGateway port.PaymentGateway,
	paymentExternalGayeway port.PaymentExternalGateway,
) port.PaymentUseCase {
	return &paymentUseCase{orderGateway, paymentGateway, paymentExternalGayeway}
}

// Create create a new payment
func (uc *paymentUseCase) Create(ctx context.Context, i dto.CreatePaymentInput) (*entity.Payment, error) {
	existentPedingPayment, err := uc.paymentGateway.GetByOrderID(ctx, i.OrderID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if existentPedingPayment.ID != 0 {
		return existentPedingPayment, nil
	}

	order, err := uc.orderGateway.FindByID(ctx, i.OrderID)
	if err != nil {
		return nil, domain.NewNotFoundError(domain.ErrOrderIsMandatory)
	}

	if len(order.OrderProducts) == 0 {
		return nil, domain.NewNotFoundError(domain.ErrOrderWithoutProducts)
	}

	paymentPayload := uc.createPaymentPayload(order)

	extPayment, err := uc.paymentExternalGateway.Create(ctx, paymentPayload)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	iPayment := &entity.Payment{
		ExternalPaymentID: extPayment.InStoreOrderID,
		OrderID:           i.OrderID,
		QrData:            extPayment.QrData,
		Status:            valueobject.PROCESSING,
	}

	payment, err := uc.paymentGateway.Create(ctx, iPayment)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	orderUpdated := &entity.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Status:     valueobject.PENDING,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  time.Now(),
	}

	if err := uc.orderGateway.Update(ctx, orderUpdated); err != nil {
		return nil, err
	}

	return payment, nil
}

func (uc *paymentUseCase) Update(ctx context.Context, payment dto.UpdatePaymentInput) (*entity.Payment, error) {
	if err := uc.paymentGateway.UpdateStatus(ctx, valueobject.CONFIRMED, payment.Resource); err != nil {
		return nil, err
	}

	paymentOUT, err := uc.paymentGateway.GetByExternalPaymentID(ctx, payment.Resource)
	if err != nil {
		return nil, err
	}

	order, err := uc.orderGateway.FindByID(ctx, paymentOUT.OrderID)
	if err != nil {
		return nil, err
	}

	orderUpdated := &entity.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Status:     valueobject.RECEIVED,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  time.Now(),
	}

	if err := uc.orderGateway.Update(ctx, orderUpdated); err != nil {
		return nil, err
	}

	return paymentOUT, nil
}

func (ps *paymentUseCase) createPaymentPayload(order *entity.Order) *entity.CreatePaymentExternalInput {
	cfg := config.LoadConfig()

	var items []entity.PaymentExternalItemsInput

	externalReference := strconv.FormatUint(order.ID, 10)

	for _, v := range order.OrderProducts {
		items = append(items, entity.PaymentExternalItemsInput{
			Title:       v.Product.Name,
			Description: v.Product.Description,
			UnitPrice:   float32(v.Product.Price),
			Category:    "marketplace",
			UnitMeasure: "unit",
			Quantity:    uint64(v.Quantity),
			TotalAmount: float32(v.Product.Price),
		})
	}

	return &entity.CreatePaymentExternalInput{
		ExternalReference: externalReference,
		TotalAmount:       order.TotalBill,
		Items:             items,
		Title:             "FIAP Tech Challenge - Product Order",
		Description:       "Purchases made at the FIAP Tech Challenge store",
		NotificationUrl:   cfg.MercadoPagoNotificationURL,
	}
}
