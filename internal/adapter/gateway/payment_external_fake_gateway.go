package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type paymentExternalFakeGayeway struct {
	dataSource port.PaymentExternalDatasource
}

func NewPaymentExternalFakeGateway(dataSource port.PaymentExternalDatasource) port.PaymentExternalGateway {
	return &paymentExternalFakeGayeway{dataSource}
}

func (g *paymentExternalFakeGayeway) Create(ctx context.Context, payment *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error) {
	return &entity.CreatePaymentExternalOutput{
		InStoreOrderID: "fake-external-payment-id",
		QrData:         "fake-qr-data",
	}, nil
}
