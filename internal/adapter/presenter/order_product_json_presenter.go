package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderProductJsonPresenter struct{}

// OrderProductJsonResponse represents the response of a orderProduct
func NewOrderProductJsonPresenter() port.OrderProductPresenter {
	return &orderProductJsonPresenter{}
}

// toOrderProductJsonResponse convert entity.OrderProduct to OrderProductJsonResponse
func toOrderProductJsonResponse(orderProduct *entity.OrderProduct) OrderProductJsonResponse {
	return OrderProductJsonResponse{
		OrderID: 	orderProduct.OrderID,
		ProductID: 	orderProduct.ProductID,
		Quantity: 	orderProduct.Quantity,
		Price: 		orderProduct.Price,
		// Order: 		toOrderJsonResponse(orderProduct.Order),
		// Product: 	toProductJsonResponse(orderProduct.Product),
		CreatedAt: orderProduct.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: orderProduct.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present write the response to the client
func (p *orderProductJsonPresenter) Present(pp dto.OrderProductPresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.OrderProduct:
		output := toOrderProductJsonResponse(v)
		pp.Writer.JSON(http.StatusOK, output)
	case []*entity.OrderProduct:
		orderProductOutputs := make([]OrderProductJsonResponse, len(v))
		for i, orderProduct := range v {
			orderProductOutputs[i] = toOrderProductJsonResponse(orderProduct)
		}

		output := &OrderProductJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			OrderProducts: orderProductOutputs,
		}
		pp.Writer.JSON(http.StatusOK, output)
	default:
		err := pp.Writer.Error(domain.NewInternalError(errors.New(domain.ErrInternalError)))
		if err != nil {
			pp.Writer.JSON(http.StatusInternalServerError, err)
		}
	}
}
