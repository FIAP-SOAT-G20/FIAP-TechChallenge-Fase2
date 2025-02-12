package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type productXmlPresenter struct{}

// NewProductXmlPresenter creates a new ProductXmlPresenter
func NewProductXmlPresenter() port.ProductPresenter {
	return &productXmlPresenter{}
}

// toResponse converts a Product entity to a ProductXmlResponse
func toXmlResponse(product *entity.Product) ProductXmlResponse {
	return ProductXmlResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present writes the response to the client
func (p *productXmlPresenter) Present(pp dto.ProductPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Product:
		output := toXmlResponse(v)
		pp.Writer.XML(http.StatusOK, output)
	case []*entity.Product:
		productOutputs := make([]ProductXmlResponse, len(v))
		for i, product := range v {
			productOutputs[i] = toXmlResponse(product)
		}

		output := &ProductXmlPaginatedResponse{
			XmlPagination: XmlPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Products: productOutputs,
		}
		pp.Writer.XML(http.StatusOK, output)
	default:
		err := pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			pp.Writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
