package dto

type CreateCustomerInput struct {
	Name   string
	Email  string
	CPF    string
	Writer ResponseWriter
}

type UpdateCustomerInput struct {
	ID     uint64
	Name   string
	Email  string
	CPF    string
	Writer ResponseWriter
}

type GetCustomerInput struct {
	ID     uint64
	Writer ResponseWriter
}

type DeleteCustomerInput struct {
	ID     uint64
	Writer ResponseWriter
}

type ListCustomersInput struct {
	Name   string
	Page   int
	Limit  int
	Writer ResponseWriter
}

type CustomerPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
	Writer ResponseWriter
}
