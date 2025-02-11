package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type productJsonPresenter struct{}

func NewProductJsonPresenter() port.ProductPresenter {
	return &productJsonPresenter{}
}

func (p *productJsonPresenter) Present(pp dto.ProductPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Product:
		output := ProductJsonResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			CategoryID:  v.CategoryID,
			CreatedAt:   v.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		pp.Writer.JSON(http.StatusOK, output)
	case []*entity.Product:
		productOutputs := make([]ProductJsonResponse, len(v))
		for i, product := range v {
			productOutputs[i] = ProductJsonResponse{
				ID:          product.ID,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				CategoryID:  product.CategoryID,
				CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			}
		}

		output := &ProductJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Products: productOutputs,
		}
		pp.Writer.JSON(http.StatusOK, output)
	default:
		err := pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			pp.Writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
