package datasource

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type FakePaymentExternalDataSource struct{}

func NewFakePaymentExternalDataSource() port.PaymentExternalDatasource {
	return &FakePaymentExternalDataSource{}
}

func (ds *FakePaymentExternalDataSource) Create(ctx context.Context, p *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error) {
	return &entity.CreatePaymentExternalOutput{
		InStoreOrderID: "fake-external-payment-id",
		QrData:         "fake-qr-data",
	}, nil
}
