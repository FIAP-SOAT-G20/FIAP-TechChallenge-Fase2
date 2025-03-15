package main

import (
	"os"

	_ "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/docs"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/gateway"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/database"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/datasource"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/httpclient"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/logger"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/route"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/server"
)

// @title						Fast Food API v2
// @version					1
// @description				### FIAP Tech Challenge Phase 2 - 10SOAT - G18
// @servers					[ { "url": "http://localhost:8080" }, { "url": "http://localhost:30001" } ]
// @BasePath					/api/v1
// @tag.name					sign-up
// @tag.description			Regiter a new customer
// @tag.name					sign-in
// @tag.description			Sign in to the system
// @tag.name					customers
// @tag.description			List, create, update and delete customers
// @tag.name					products
// @tag.description			List, create, update and delete products
// @tag.name					orders
// @tag.description			List, create, update and delete orders
// @tag.name					payments
// @tag.description			Process payments
// @tag.name					staffs
// @tag.description			List, create, update and delete staff
// @tag.name					health-check
// @tag.description			Health check
//
// @externalDocs.description	GitHub Repository
// @externalDocs.url			https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	cfg := config.LoadConfig()

	loggerInstance := logger.NewLogger(cfg)

	db, err := database.NewPostgresConnection(cfg, loggerInstance.Logger)
	if err != nil {
		loggerInstance.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	if err := db.Migrate(); err != nil {
		loggerInstance.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}

	httpClient := httpclient.NewRestyClient(cfg, loggerInstance.Logger)

	handlers := setupHandlers(db, httpClient, cfg)

	srv := server.NewServer(cfg, loggerInstance, handlers)
	if err := srv.Start(); err != nil {
		loggerInstance.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

func setupHandlers(db *database.Database, httpClient *httpclient.HTTPClient, cfg *config.Config) *route.Handlers {
	// Datasources
	productDS := datasource.NewProductDataSource(db.DB)
	customerDS := datasource.NewCustomerDataSource(db.DB)
	orderDS := datasource.NewOrderDataSource(db.DB)
	orderProductDS := datasource.NewOrderProductDataSource(db.DB)
	staffDS := datasource.NewStaffDataSource(db.DB)
	orderHistoryDS := datasource.NewOrderHistoryDataSource(db.DB)
	paymentDS := datasource.NewPaymentDataSource(db.DB)
	// paymentExternalDS := datasource.NewPaymentExternalDataSource(httpClient.Client) // Mercado Pago
	paymentExternalDS := datasource.NewFakePaymentExternalDataSource(httpClient, cfg) // Fake Mercado Pago

	// Gateways
	productGateway := gateway.NewProductGateway(productDS)
	customerGateway := gateway.NewCustomerGateway(customerDS)
	orderHistoryGateway := gateway.NewOrderHistoryGateway(orderHistoryDS)
	orderGateway := gateway.NewOrderGateway(orderDS)
	orderProductGateway := gateway.NewOrderProductGateway(orderProductDS)
	staffGateway := gateway.NewStaffGateway(staffDS)
	paymentGateway := gateway.NewPaymentGateway(paymentDS, paymentExternalDS)

	// Use cases
	productUC := usecase.NewProductUseCase(productGateway)
	customerUC := usecase.NewCustomerUseCase(customerGateway)
	orderHistoryUC := usecase.NewOrderHistoryUseCase(orderHistoryGateway)
	orderUC := usecase.NewOrderUseCase(orderGateway, orderHistoryUC)
	orderProductUC := usecase.NewOrderProductUseCase(orderProductGateway)
	staffUC := usecase.NewStaffUseCase(staffGateway)
	paymentUC := usecase.NewPaymentUseCase(paymentGateway, orderUC)

	// Controllers
	productController := controller.NewProductController(productUC)
	customerController := controller.NewCustomerController(customerUC)
	orderController := controller.NewOrderController(orderUC)
	orderProductController := controller.NewOrderProductController(orderProductUC)
	staffController := controller.NewStaffController(staffUC)
	orderHistoryController := controller.NewOrderHistoryController(orderHistoryUC)
	paymentController := controller.NewPaymentController(paymentUC)

	// Handlers
	productHandler := handler.NewProductHandler(productController)
	customerHandler := handler.NewCustomerHandler(customerController)
	orderHandler := handler.NewOrderHandler(orderController)
	orderProductHandler := handler.NewOrderProductHandler(orderProductController)
	staffHandler := handler.NewStaffHandler(staffController)
	healthCheckHandler := handler.NewHealthCheckHandler()
	orderHistoryHandler := handler.NewOrderHistoryHandler(orderHistoryController)
	paymentHandler := handler.NewPaymentHandler(paymentController)

	return &route.Handlers{
		Product:      productHandler,
		Customer:     customerHandler,
		Order:        orderHandler,
		OrderProduct: orderProductHandler,
		OrderHistory: orderHistoryHandler,
		Staff:        staffHandler,
		HealthCheck:  healthCheckHandler,
		Payment:      paymentHandler,
	}
}
