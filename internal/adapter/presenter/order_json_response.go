package presenter

type OrderJsonResponse struct {
	ID         uint64                 `json:"id"`
	CustomerID uint64                 `json:"customer_id"`
	TotalBill  float32                `json:"total_bill"`
	Status     string                 `json:"status"`
	Products   []ProductsJsonResponse `json:"products,omitempty"`
	CreatedAt  string                 `json:"created_at" example:"2024-02-09T10:00:00Z"`
	UpdatedAt  string                 `json:"updated_at" example:"2024-02-09T10:00:00Z"`
}

type OrderJsonPaginatedResponse struct {
	JsonPagination
	Orders []OrderJsonResponse `json:"orders"`
}

type ProductsJsonResponse struct {
	ProductJsonResponse
	Quantity uint32 `json:"quantity"`
}
