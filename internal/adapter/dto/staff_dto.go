package dto

type CreateStaffInput struct {
	Name string
	Role string
}

type UpdateStaffInput struct {
	ID   uint64
	Name string
	Role string
}

type GetStaffInput struct {
	ID uint64
}

type DeleteStaffInput struct {
	ID uint64
}

type ListStaffsInput struct {
	Name  string
	Role  string
	Page  int
	Limit int
}

type StaffPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
}
