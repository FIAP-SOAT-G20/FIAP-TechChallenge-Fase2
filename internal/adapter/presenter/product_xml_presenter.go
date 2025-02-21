package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type productXmlPresenter struct {
	writer ResponseWriter
}

// NewProductXmlPresenter creates a new ProductXmlPresenter
func NewProductXmlPresenter(writer ResponseWriter) port.Presenter {
	return &productXmlPresenter{writer}
}

// Present writes the response to the client
func (p *productXmlPresenter) Present(pp dto.PresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.Product:
		output := toProductXmlResponse(v)
		p.writer.XML(http.StatusOK, output)
	case []*entity.Product:
		productOutputs := make([]ProductXmlResponse, len(v))
		for i, product := range v {
			productOutputs[i] = toProductXmlResponse(product)
		}

		output := &ProductXmlPaginatedResponse{
			XmlPagination: XmlPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Products: productOutputs,
		}
		p.writer.XML(http.StatusOK, output)
	default:
		p.writer.XML(
			http.StatusInternalServerError,
			domain.NewInternalError(errors.New(domain.ErrInternalError)),
		)
	}
}

// toProductXmlResponse converts a Product entity to a ProductXmlResponse
func toProductXmlResponse(product *entity.Product) ProductXmlResponse {
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
