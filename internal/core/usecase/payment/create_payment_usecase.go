package payment

import (
	"context"
	"os"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createPaymentUseCase struct {
	paymentGateway  port.PaymentGateway
	orderGateway    port.OrderGateway
	paymentExternal port.PaymentExternalDatasource
	presenter       port.PaymentPresenter
}

// NewCreatePaymentUseCase creates a new createPaymentUseCase
func NewCreatePaymentUseCase(paymentGateway port.PaymentGateway, orderGateway port.OrderGateway, paymentExternal port.PaymentExternalDatasource, presenter port.PaymentPresenter) port.CreatePaymentUseCase {
	return &createPaymentUseCase{
		paymentGateway:  paymentGateway,
		orderGateway:    orderGateway,
		paymentExternal: paymentExternal,
		presenter:       presenter,
	}
}

// Execute create a new payment
func (uc *createPaymentUseCase) Execute(ctx context.Context, OrderID uint64, writer dto.ResponseWriter) error {
	existentPedingPayment, err := uc.paymentGateway.GetPaymentByOrderIDAndStatus(ctx, valueobject.PROCESSING, OrderID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if existentPedingPayment.ID != 0 {

		uc.presenter.Present(dto.PaymentPresenterInput{
			Result: existentPedingPayment,
			Writer: writer,
		})
		return nil
	}

	order, err := uc.orderGateway.FindByID(ctx, OrderID)
	if err != nil {
		return domain.NewNotFoundError("order not found")
	}

	if len(order.OrderProducts) == 0 {
		return domain.NewNotFoundError("no products")
	}

	paymentPayload := uc.createPaymentPayload(order)

	extPayment, err := uc.paymentExternal.CreatePayment(paymentPayload)
	if err != nil {
		return domain.NewInternalError(err)
	}

	iPayment := &entity.Payment{
		Status:            valueobject.PROCESSING,
		ExternalPaymentID: extPayment.InStoreOrderID,
		OrderID:           OrderID,
		QrData:            extPayment.QrData,
	}

	payment, err := uc.paymentGateway.Create(ctx, iPayment)
	if err != nil {
		return nil
	}

	order.Status = valueobject.PENDING
	if err := uc.orderGateway.Update(ctx, order); err != nil {
		return err
	}

	uc.presenter.Present(dto.PaymentPresenterInput{
		Writer: writer,
		Result: payment,
	})

	return nil
}

func (ps *createPaymentUseCase) createPaymentPayload(order *entity.Order) *entity.CreatePaymentIN {
	var items []entity.ItemsIN

	externalReference := strconv.FormatUint(order.ID, 10)

	for _, v := range order.OrderProducts {
		items = append(items, entity.ItemsIN{
			Title:       v.Product.Name,
			Description: v.Product.Description,
			UnitPrice:   float32(v.Product.Price),
			Category:    "marketplace",
			UnitMeasure: "unit",
			Quantity:    uint64(v.Quantity),
			TotalAmount: float32(v.Product.Price),
		})
	}

	return &entity.CreatePaymentIN{
		ExternalReference: externalReference,
		TotalAmount:       order.TotalBill,
		Items:             items,
		Title:             "FIAP Tech Challenge - Product Order",
		Description:       "Purchases made at the FIAP Tech Challenge store",
		NotificationUrl:   os.Getenv("MERCADO_PAGO_NOTIFICATION_URL"),
	}
}
