package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type paymentGayeway struct {
	dataSource port.PaymentDataSource
}

func NewPaymentGateway(dataSource port.PaymentDataSource) port.PaymentGateway {
	return &paymentGayeway{dataSource}
}

func (g *paymentGayeway) GetPaymentByOrderIDAndStatus(ctx context.Context, status valueobject.PaymentStatus, orderID uint64) (*entity.Payment, error) {
	return g.dataSource.GetPaymentByOrderIDAndStatus(ctx, status, orderID)
}

func (g *paymentGayeway) Create(ctx context.Context, p *entity.Payment) (*entity.Payment, error) {
	return g.dataSource.Create(ctx, p)
}
