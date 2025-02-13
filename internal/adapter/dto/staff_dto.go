package dto

type CreateStaffInput struct {
	Name   string
	Role   string
	Writer ResponseWriter
}

type UpdateStaffInput struct {
	ID     uint64
	Name   string
	Role   string
	Writer ResponseWriter
}

type GetStaffInput struct {
	ID     uint64
	Writer ResponseWriter
}

type DeleteStaffInput struct {
	ID     uint64
	Writer ResponseWriter
}

type ListStaffsInput struct {
	Name   string
	Role   string
	Page   int
	Limit  int
	Writer ResponseWriter
}

type StaffPresenterInput struct {
	Result any
	Total  int64
	Page   int
	Limit  int
	Writer ResponseWriter
}
