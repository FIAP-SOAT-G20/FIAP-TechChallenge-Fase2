package dto

type ProductRequest struct {
	Name        string  `json:"name" example:"Produto A"`
	Description string  `json:"description" example:"Descrição do Produto A"`
	Price       float64 `json:"price" example:"99.99"`
	CategoryID  uint64  `json:"category_id" example:"1"`
}

type ProductResponse struct {
	ID          uint64  `json:"id" example:"1"`
	Name        string  `json:"name" example:"Produto A"`
	Description string  `json:"description" example:"Descrição do Produto A"`
	Price       float64 `json:"price" example:"99.99"`
	CategoryID  uint64  `json:"category_id" example:"1"`
	CreatedAt   string  `json:"created_at" example:"2024-02-09T10:00:00Z"`
	UpdatedAt   string  `json:"updated_at" example:"2024-02-09T10:00:00Z"`
}

type ProductListRequest struct {
	Name       string `json:"name" example:"Produto"`
	CategoryID uint64 `json:"category_id" example:"1"`
	Page       int    `json:"page" example:"1"`
	Limit      int    `json:"limit" example:"10"`
}

type PaginatedResponse struct {
	Pagination
	Products []ProductResponse `json:"products"`
}
