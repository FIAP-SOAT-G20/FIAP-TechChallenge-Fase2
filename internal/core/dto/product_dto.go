package dto

type CreateProductInput struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
}

type UpdateProductInput struct {
	ID          uint64
	Name        string
	Description string
	Price       float64
	CategoryID  uint64
}

type GetProductInput struct {
	ID uint64
}

type DeleteProductInput struct {
	ID uint64
}

type ListProductsInput struct {
	Name       string
	CategoryID uint64
	Page       int
	Limit      int
}
