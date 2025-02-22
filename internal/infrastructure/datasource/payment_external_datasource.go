package datasource

import (
	"fmt"
	"os"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/go-resty/resty/v2"
)

type PaymentExternalDataSource struct {
}

func NewPaymentExternal() port.PaymentExternalDatasource {
	return &PaymentExternalDataSource{}
}

func (ps *PaymentExternalDataSource) CreatePayment(payment *entity.CreatePaymentIN) (*entity.CreatePaymentOUT, error) {
	body := entity.NewPaymentRequest(payment)

	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("MERCADO_PAGO_TOKEN"))).
		SetHeader("Content-Type", "application/json")

	resp, err := client.R().
		SetBody(body).
		SetResult(&entity.CreatePaymentResponse{}).
		Post(os.Getenv("MERCADO_PAGO_URL"))
	if err != nil {
		return nil, fmt.Errorf("error to create payment: %w", err)
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("error: response status %d", resp.StatusCode())
	}

	response := entity.ToCreatePaymentOUTDomain(resp.Result().(*entity.CreatePaymentResponse))

	return response, nil
}
