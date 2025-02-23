package port

import (
	"context"
)

type PaymentController interface {
	Create(ctx context.Context, presenter Presenter, OrderID uint64) error
}
