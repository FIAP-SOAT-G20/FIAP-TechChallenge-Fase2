package external

import (
	"fmt"
	"os"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type PaymentExternal struct {
}

func NewPaymentExternal() *PaymentExternal {
	return &PaymentExternal{}
}

func (ps *PaymentExternal) CreatePayment(payment *entity.CreatePaymentIN) (*entity.CreatePaymentOUT, error) {
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

func (ps *PaymentExternal) CreatePaymentMock(payment *entity.CreatePaymentIN) (*entity.CreatePaymentOUT, error) {
	return &entity.CreatePaymentOUT{
		InStoreOrderID: uuid.New().String(),
		QrData:         "https://www.fiap-10-soat-g20.com.br/qr/123456",
	}, nil
}
