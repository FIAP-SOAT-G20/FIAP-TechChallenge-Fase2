package http

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/go-resty/resty/v2"
)

func NewRestyClient(cfg *config.Config, logger *slog.Logger) *resty.Client {
	httpCLient := resty.New().
		SetTimeout(10*time.Second). // TODO: set timeout in config (ENV)
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.MercadoPagoToken)).
		SetHeader("Content-Type", "application/json")

	logger.Info("resty client created")

	return httpCLient
}
