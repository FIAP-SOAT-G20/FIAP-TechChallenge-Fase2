package presenter

type OrderProductJsonResponse struct {
	OrderID   uint64              `json:"order_id"`
	ProductID uint64              `json:"product_id"`
	Price     float32             `json:"price"`
	Quantity  uint32              `json:"quantity"`
	Order     OrderJsonResponse   `json:"order"`
	Product   ProductJsonResponse `json:"product"`
	CreatedAt string              `json:"created_at" example:"2024-02-09T10:00:00Z"`
	UpdatedAt string              `json:"updated_at" example:"2024-02-09T10:00:00Z"`
}

func NewOrderProductJsonResponse(orderID uint64, productID uint64, price float32, quantity uint32) *OrderProductJsonResponse {
	orderProduct := &OrderProductJsonResponse{
		OrderID:   orderID,
		ProductID: productID,
		Price:     price,
		Quantity:  quantity,
	}

	return orderProduct
}

type OrderProductJsonPaginatedResponse struct {
	JsonPagination
	OrderProducts []OrderProductJsonResponse `json:"order_products"`
}
