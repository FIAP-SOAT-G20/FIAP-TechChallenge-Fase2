package dto

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

type ProductPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
	Writer ResponseWriter
}
