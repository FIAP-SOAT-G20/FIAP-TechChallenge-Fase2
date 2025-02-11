package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type productXmlPresenter struct{}

func NewProductPresenter() port.ProductPresenter {
	return &productXmlPresenter{}
}

func (p *productXmlPresenter) Present(pp port.ProductPresenterDTO) {
	switch v := pp.Result.(type) {
	case *entity.Product:
		output := dto.ProductXmlResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			CategoryID:  v.CategoryID,
			CreatedAt:   v.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   v.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		pp.Writer.XML(http.StatusOK, output)
	case []*entity.Product:
		productOutputs := make([]dto.ProductXmlResponse, len(v))
		for i, product := range v {
			productOutputs[i] = dto.ProductXmlResponse{
				ID:          product.ID,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				CategoryID:  product.CategoryID,
				CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			}
		}

		output := &dto.ProductXmlPaginatedResponse{
			XmlPagination: dto.XmlPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Products: productOutputs,
		}
		pp.Writer.XML(http.StatusOK, output)
	default:
		pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
	}

}
