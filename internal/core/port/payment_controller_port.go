package port

import (
	"context"
)

type PaymentController interface {
	CreatePayment(ctx context.Context, presenter Presenter, OrderID uint64) error
}
