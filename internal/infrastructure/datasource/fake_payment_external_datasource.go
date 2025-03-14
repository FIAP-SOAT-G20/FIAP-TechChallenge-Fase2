package datasource

import (
	"context"
	"encoding/json"
	"errors"

	datasource_request "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/datasource/request"
	datasource_response "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/datasource/response"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/httpclient"
)

type FakePaymentExternalDataSource struct {
	client *httpclient.HTTPClient
	cfg    *config.Config
}

func NewFakePaymentExternalDataSource(client *httpclient.HTTPClient, cfg *config.Config) port.PaymentExternalDatasource {
	return &FakePaymentExternalDataSource{client: client, cfg: cfg}
}

func (ds *FakePaymentExternalDataSource) Create(ctx context.Context, p *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error) {
	fakeMercadoPagoRequest := datasource_request.NewFakeMercadoPagoRequest(p)
	fakeMercadoPagoResponse := datasource_response.MercadoPagoResponse{}

	response, err := ds.client.NewRequest().
		SetHeader("Content-Type", "application/json").
		SetBody(fakeMercadoPagoRequest).
		SetResult(&fakeMercadoPagoResponse).
		Post(ds.cfg.FakeMercadoPagoURL)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, errors.New(domain.ErrFailedToCreatePaymentExternal)
	}

	err = json.Unmarshal(response.Body(), &fakeMercadoPagoResponse)
	if err != nil {
		return nil, err
	}

	return datasource_response.NewCreatePaymentExternalOutput(&fakeMercadoPagoResponse), nil
}
