package route

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/middleware"
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

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Router{
		engine: engine,
		logger: logger,
	}
}

// RegisterRoutes configure all routes of the application
func (r *Router) RegisterRoutes(handlers *Handlers) {
	// API v1
	v1 := r.engine.Group("/api/v1")
	{
		handlers.Product.Register(v1.Group("/products"))
		handlers.Customer.Register(v1.Group("/customers"))
		handlers.Staff.Register(v1.Group("/staffs"))

		// Orders
		orders := v1.Group("/orders")
		handlers.Order.Register(orders)

		// Order Products
		orderProducts := v1.Group("/orders/products")
		handlers.OrderProduct.Register(orderProducts)

		// Payments
		payments := v1.Group("/payments")
		handlers.Payment.Register(payments)

		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "UP"})
		})
	}
	// Health check
	handlers.HealthCheck.Register(r.engine.Group("/health"))
}

// Engine returns the gin engine
func (r *Router) Engine() *gin.Engine {
	return r.engine
}

// Handlers contains all handlers of the application
type Handlers struct {
	Product      *handler.ProductHandler
	Customer     *handler.CustomerHandler
	Staff        *handler.StaffHandler
	Order        *handler.OrderHandler
	OrderProduct *handler.OrderProductHandler
	HealthCheck  *handler.HealthCheckHandler
	Payment      *handler.PaymentHandler
}
