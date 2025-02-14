package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type productJsonPresenter struct{}

// ProductJsonResponse represents the response of a product
func NewProductJsonPresenter() port.ProductPresenter {
	return &productJsonPresenter{}
}

// toProductJsonResponse convert entity.Product to ProductJsonResponse
func toProductJsonResponse(product *entity.Product) ProductJsonResponse {
	return ProductJsonResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present write the response to the client
func (p *productJsonPresenter) Present(pp dto.ProductPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Product:
		output := toProductJsonResponse(v)
		pp.Writer.JSON(http.StatusOK, output)
	case []*entity.Product:
		productOutputs := make([]ProductJsonResponse, len(v))
		for i, product := range v {
			productOutputs[i] = toProductJsonResponse(product)
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
