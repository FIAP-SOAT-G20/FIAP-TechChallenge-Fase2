package dto

type ProductRequest struct {
	Name        string  `json:"name" example:"Produto A"`
	Description string  `json:"description" example:"Descrição do Produto A"`
	Price       float64 `json:"price" example:"99.99"`
	CategoryID  uint64  `json:"category_id" example:"1"`
}

type ProductListRequest struct {
	Name       string `json:"name" example:"Produto"`
	CategoryID uint64 `json:"category_id" example:"1"`
	Page       int    `json:"page" example:"1"`
	Limit      int    `json:"limit" example:"10"`
}

type ProductJsonResponse struct {
	ID          uint64  `json:"id" example:"1"`
	Name        string  `json:"name" example:"Produto A"`
	Description string  `json:"description" example:"Descrição do Produto A"`
	Price       float64 `json:"price" example:"99.99"`
	CategoryID  uint64  `json:"category_id" example:"1"`
	CreatedAt   string  `json:"created_at" example:"2024-02-09T10:00:00Z"`
	UpdatedAt   string  `json:"updated_at" example:"2024-02-09T10:00:00Z"`
}

type ProductJsonPaginatedResponse struct {
	JsonPagination
	Products []ProductJsonResponse `json:"products"`
}

type ProductXmlResponse struct {
	ID          uint64  `xml:"id" example:"1"`
	Name        string  `xml:"name" example:"Produto A"`
	Description string  `xml:"description" example:"Descrição do Produto A"`
	Price       float64 `xml:"price" example:"99.99"`
	CategoryID  uint64  `xml:"category_id" example:"1"`
	CreatedAt   string  `xml:"created_at" example:"2024-02-09T10:00:00Z"`
	UpdatedAt   string  `xml:"updated_at" example:"2024-02-09T10:00:00Z"`
}

type ProductXmlPaginatedResponse struct {
	XmlPagination
	Products []ProductXmlResponse `xml:"products"`
}

type CreateProductInput struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
	Writer      ResponseWriter
}

type UpdateProductInput struct {
	ID          uint64
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
	Writer      ResponseWriter
}

type GetProductInput struct {
	ID     uint64
	Writer ResponseWriter
}

type DeleteProductInput struct {
	ID     uint64
	Writer ResponseWriter
}

type ListProductsInput struct {
	Name       string
	CategoryID uint64
	Page       int
	Limit      int
	Writer     ResponseWriter
}

type ProductOutput struct {
	ID          uint64
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
	CreatedAt   string
	UpdatedAt   string
}
