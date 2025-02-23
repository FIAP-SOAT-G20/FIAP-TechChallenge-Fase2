package httpclient

import (
	"fmt"
	"log/slog"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/go-resty/resty/v2"
)

type HTTPClient struct {
	*resty.Client
}

func NewRestyClient(cfg *config.Config, logger *slog.Logger) *HTTPClient {
	httpCLient := resty.New().
		SetTimeout(cfg.MercadoPagoTimeout).
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.MercadoPagoToken)).
		SetHeader("Content-Type", "application/json")

	logger.Info("resty client created")

	return &HTTPClient{httpCLient}
}
