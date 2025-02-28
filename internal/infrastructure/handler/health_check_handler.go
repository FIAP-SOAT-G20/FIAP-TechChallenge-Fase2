package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/response"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Register(router *gin.RouterGroup) {
	router.GET("", h.HealthCheck)
	router.GET("/", h.HealthCheck)
}

// HealthCheck godoc
//
//	@Summary		Application HealthCheck
//	@Description	Checks application health
//	@Tags			health-check
//	@Produce		json
//	@Success		200	{object}	response.HealthCheckResponse
//	@Failure		500	{object}	string							"Internal server error"
//	@Failure		503	{object}	response.HealthCheckResponse	"Service Unavailable"
//	@Router			/health [GET]
func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	cfg := config.LoadConfig()
	hc := &response.HealthCheckResponse{
		Status: response.HealthCheckStatusPass,
		Checks: map[string]response.HealthCheckVerifications{
			"postgres:status": {
				ComponentId: "db:postgres",
				Status:      response.HealthCheckStatusPass,
				Time:        time.Now(),
			},
		},
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if db.Ping() != nil {
		hc.Status = response.HealthCheckStatusFail
		hc.Checks["postgres:status"] = response.HealthCheckVerifications{
			ComponentId: "db:postgres",
			Status:      response.HealthCheckStatusFail,
			Time:        time.Now(),
		}
		c.JSON(http.StatusServiceUnavailable, hc)
		return
	}

	switch hc.Status {
	case response.HealthCheckStatusFail:
		c.JSON(http.StatusServiceUnavailable, hc)
	default:
		c.JSON(http.StatusOK, hc)
	}
}
