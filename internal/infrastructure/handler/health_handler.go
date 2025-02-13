package handler

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GroupRouterPattern() string {
	return "/health"
}

func (h *HealthHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.Health)
}

// HealthCheck Status da aplicação
//
//	@Summary		Health Check
//	@Description	Endpoint para verificar a saúde da aplicação
//	@Tags			health
//	@Produce		json
//	@Success		200			{object}	string
//	@Failure		500			{object}	string
//	@Router			/health [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}
