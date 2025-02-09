package dto

// ProductRequest representa a requisição para criar um produto
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
	Total    int64             `json:"total" example:"100"`
	Page     int               `json:"page" example:"1"`
	Limit    int               `json:"limit" example:"10"`
	Products []ProductResponse `json:"products"`
}

type ErrorResponse struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Erro de validação"`
	Errors  interface{} `json:"errors,omitempty"`
}
