package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"

type ProductPresenter interface {
	Present(pp dto.ProductPresenterInput)
}

type ProductPresenterDTO struct {
	Result any
	Total  int64
	Page   int
	Limit  int
	Writer dto.ResponseWriter
}
