package request

type CreatePaymentUriRequest struct {
	OrderID uint64 `uri:"order_id" binding:"required"`
}
