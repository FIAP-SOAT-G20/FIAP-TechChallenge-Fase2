package routes

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/docs"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/middleware"
)

type Router struct {
	engine *gin.Engine
	logger *slog.Logger
}

type IRouter interface {
	GroupRouterPattern() string
	Register(router *gin.RouterGroup)
}

func InitGinEngine(cfg *config.Config) {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func NewRouter(logger *slog.Logger, cfg *config.Config) *Router {
	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	// Global middlewares
	engine.Use(
		middleware.RequestID(),
		middleware.Logger(logger),
		middleware.ErrorHandler(logger),
		middleware.Recovery(logger),
		middleware.CORS(),
	)

	docs.SwaggerInfo.BasePath = "/api/v1"
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Router{
		engine: engine,
		logger: logger,
	}
}

// RegisterRoutes configura todas as rotas da aplicação
func RegisterRoutes(r *Router, handlers []IRouter) {
	// API v1
	v1 := r.engine.Group("/api/v1")
	{
		for _, handler := range handlers {
			handler.Register(v1.Group(handler.GroupRouterPattern()))
		}
	}
}

// Engine retorna o gin.Engine para ser usado no servidor HTTP
func (r *Router) Engine() *gin.Engine {
	return r.engine
}
