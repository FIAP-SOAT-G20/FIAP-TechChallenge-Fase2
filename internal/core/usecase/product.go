package usecase

type CreateProductInput struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
}

type UpdateProductInput struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
}

type ListProductsInput struct {
	Name       string
	CategoryID uint64
	Page       int
	Limit      int
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

type ListProductPaginatedOutput struct {
	PaginatedOutput
	Products []ProductOutput
}
