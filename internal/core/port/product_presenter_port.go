package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"

type ProductPresenter interface {
	Present(pp dto.ProductPresenterInput)
}

type ProductPresenterDTO struct {
	Writer dto.ResponseWriter
	Result any
	Total  int64
	Page   int
	Limit  int
}
