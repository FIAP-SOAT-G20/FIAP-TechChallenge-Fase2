package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"

type StaffPresenter interface {
	Present(pp dto.StaffPresenterInput)
}

type StaffPresenterDTO struct {
	Writer dto.ResponseWriter
	Result any
	Total  int64
	Page   int
	Limit  int
}
