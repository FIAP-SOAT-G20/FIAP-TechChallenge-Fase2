package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type paymentGayeway struct {
	dataSource port.PaymentDataSource
}

func NewPaymentGateway(dataSource port.PaymentDataSource) port.PaymentGateway {
	return &paymentGayeway{
		dataSource: dataSource,
	}
}

func (pg *paymentGayeway) GetPaymentByOrderIDAndStatus(ctx context.Context, status entity.PaymentStatus, orderID uint64) (*entity.Payment, error) {
	return pg.dataSource.GetPaymentByOrderIDAndStatus(ctx, status, orderID)
}

func (pg *paymentGayeway) Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	return pg.dataSource.Create(ctx, payment)
}
