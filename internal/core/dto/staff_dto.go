package dto

import valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"

type CreateStaffInput struct {
	Name string
	Role valueobject.StaffRole
}

type UpdateStaffInput struct {
	ID   uint64
	Name string
	Role valueobject.StaffRole
}

type GetStaffInput struct {
	ID uint64
}

type DeleteStaffInput struct {
	ID uint64
}

type ListStaffsInput struct {
	Name  string
	Role  valueobject.StaffRole
	Page  int
	Limit int
}
