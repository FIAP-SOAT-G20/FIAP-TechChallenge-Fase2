package routes

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"tech-challenge-2-app-example/docs"
	"tech-challenge-2-app-example/internal/adapters/handler"
	"tech-challenge-2-app-example/internal/infrastructure/middleware"
)

type Router struct {
	engine *gin.Engine
	logger *slog.Logger
}

func NewRouter(logger *slog.Logger, environment string) *Router {
	// Set Gin mode
	if environment == "production" {
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

	docs.SwaggerInfo.BasePath = ""
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Router{
		engine: engine,
		logger: logger,
	}
}

// RegisterRoutes configura todas as rotas da aplicação
func (r *Router) RegisterRoutes(handlers *Handlers) {
	// API v1
	v1 := r.engine.Group("/api/v1")
	{
		// Products
		products := v1.Group("/products")
		handlers.Product.Register(products)

		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "UP"})
		})

		// Adicione outras rotas aqui
	}
}

// Engine retorna o gin.Engine para ser usado no servidor HTTP
func (r *Router) Engine() *gin.Engine {
	return r.engine
}

// Handlers agrupa todos os handlers da aplicação
type Handlers struct {
	Product *handler.ProductHandler
	// Adicione outros handlers aqui
}
