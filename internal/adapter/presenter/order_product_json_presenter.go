package presenter

import (
	"errors"
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type orderProductJsonPresenter struct {
	writer ResponseWriter
}

// OrderProductJsonResponse represents the response of a orderProduct
func NewOrderProductJsonPresenter(writer ResponseWriter) port.Presenter {
	return &orderProductJsonPresenter{writer}
}

// Present write the response to the client
func (p *orderProductJsonPresenter) Present(pp dto.PresenterInput) {
	switch v := pp.Result.(type) {
	case *entity.OrderProduct:
		output := ToOrderProductJsonResponse(v)
		p.writer.JSON(http.StatusOK, output)
	case []*entity.OrderProduct:
		orderProductOutputs := make([]OrderProductJsonResponse, len(v))
		for i, orderProduct := range v {
			orderProductOutputs[i] = ToOrderProductJsonResponse(orderProduct)
		}

		output := &OrderProductJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			OrderProducts: orderProductOutputs,
		}
		p.writer.JSON(http.StatusOK, output)
	default:
		p.writer.JSON(
			http.StatusInternalServerError,
			domain.NewInternalError(errors.New(domain.ErrInternalError)),
		)
	}
}

// ToOrderProductJsonResponse convert entity.OrderProduct to OrderProductJsonResponse
func ToOrderProductJsonResponse(orderProduct *entity.OrderProduct) OrderProductJsonResponse {
	order := ToOrderJsonResponse(&orderProduct.Order)
	order.TotalBill = ""
	return OrderProductJsonResponse{
		OrderID:   orderProduct.OrderID,
		ProductID: orderProduct.ProductID,
		Quantity:  orderProduct.Quantity,
		Order:     order,
		Product:   ToProductJsonResponse(&orderProduct.Product),
		CreatedAt: orderProduct.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: orderProduct.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
