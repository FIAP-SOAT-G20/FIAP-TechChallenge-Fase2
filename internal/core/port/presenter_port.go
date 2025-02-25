package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"

type Presenter interface {
	Present(dto.PresenterInput) ([]byte, error)
}
