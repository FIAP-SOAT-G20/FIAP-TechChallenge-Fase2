package datasource

import (
	"fmt"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/response"
	"github.com/go-resty/resty/v2"
)

type PaymentExternalDataSource struct {
}

func NewPaymentExternal() port.PaymentExternalDatasource {
	return &PaymentExternalDataSource{}
}

func (ps *PaymentExternalDataSource) CreatePayment(payment *entity.CreatePaymentIN) (*entity.CreatePaymentOUT, error) {
	cfg := config.LoadConfig()

	body := request.NewPaymentRequest(payment)

	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.MercadoPagoToken)).
		SetHeader("Content-Type", "application/json")

	resp, err := client.R().
		SetBody(body).
		SetResult(&response.CreatePaymentResponse{}).
		Post(cfg.MercadoPagoURL)
	if err != nil {
		return nil, fmt.Errorf("error to create payment: %w", err)
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("error: response status %d", resp.StatusCode())
	}

	response := response.ToCreatePaymentOUTDomain(resp.Result().(*response.CreatePaymentResponse))

	return response, nil
}
