package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type paymentExternalGayeway struct {
	dataSource port.PaymentExternalDatasource
}

func NewPaymentExternalGateway(dataSource port.PaymentExternalDatasource) port.PaymentExternalGateway {
	return &paymentExternalGayeway{dataSource}
}

func (g *paymentExternalGayeway) Create(ctx context.Context, payment *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error) {
	return g.dataSource.Create(ctx, payment)
}
