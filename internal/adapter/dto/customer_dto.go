package dto

type CreateCustomerInput struct {
	Name  string
	Email string
	CPF   string
}

type UpdateCustomerInput struct {
	ID    uint64
	Name  string
	Email string
}

type GetCustomerInput struct {
	ID uint64
}

type DeleteCustomerInput struct {
	ID uint64
}

type ListCustomersInput struct {
	Name  string
	Page  int
	Limit int
}

type CustomerPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
}
