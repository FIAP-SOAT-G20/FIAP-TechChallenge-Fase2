package main

import (
	"os"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/gateway"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/customer"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/product"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/database"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/datasource"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/logger"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/route"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/server"
)

// @title						FIAP Tech Challenge Fase 2 - 10SOAT - G80
// @version					1
// @description				### API de um Fast Food para o Tech Challenge da FIAP - Fase 2 - 10SOAT - G18
// @servers					[ { "url": "http://localhost:8080" } ]
// @host						localhost:8080
// @BasePath					/api/v1
// @tag.name					sign-up
// @tag.description			Regiter a new customer or staff
// @tag.name					products
// @tag.description			List, create, update and delete products
// @tag.name					payments
// @tag.description			Process payments
// @tag.name					sign-in
// @tag.description			Sign in to the system
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

	handlers := setupHandlers(db)

	srv := server.NewServer(cfg, loggerInstance.Logger, handlers)
	if err := srv.Start(); err != nil {
		loggerInstance.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

func setupHandlers(db *database.Database) *route.Handlers {
	// Datasource's
	productDS := datasource.NewProductDataSource(db.DB)
	customerDS := datasource.NewCustomerDataSource(db.DB)
	orderDS := datasource.NewOrderDataSource(db.DB)

	// Gateways
	productGateway := gateway.NewProductGateway(productDS)
	customerGateway := gateway.NewCustomerGateway(customerDS)
	orderGateway := gateway.NewOrderGateway(orderDS)

	// Presenters
	productPresenter := presenter.NewProductJsonPresenter()
	// productPresenter := presenter.NewProductXmlPresenter()
	customerPresenter := presenter.NewCustomerJsonPresenter()
	orderPresenter := presenter.NewOrderJsonPresenter()

	// Use cases
	listProductsUC := product.NewListProductsUseCase(productGateway, productPresenter)
	createProductUC := product.NewCreateProductUseCase(productGateway, productPresenter)
	getProductUC := product.NewGetProductUseCase(productGateway, productPresenter)
	updateProductUC := product.NewUpdateProductUseCase(productGateway, productPresenter)
	deleteProductUC := product.NewDeleteProductUseCase(productGateway, productPresenter)
	listCustomersUC := customer.NewListCustomersUseCase(customerGateway, customerPresenter)
	createCustomerUC := customer.NewCreateCustomerUseCase(customerGateway, customerPresenter)
	getCustomerUC := customer.NewGetCustomerUseCase(customerGateway, customerPresenter)
	updateCustomerUC := customer.NewUpdateCustomerUseCase(customerGateway, customerPresenter)
	deleteCustomerUC := customer.NewDeleteCustomerUseCase(customerGateway, customerPresenter)
	createOrderUC := order.NewCreateOrderUseCase(orderGateway, orderPresenter)
	listOrdersUC := order.NewListOrdersUseCase(orderGateway, orderPresenter)

	// Controllers
	productController := controller.NewProductController(
		listProductsUC,
		createProductUC,
		getProductUC,
		updateProductUC,
		deleteProductUC,
	)
	customerController := controller.NewCustomerController(
		listCustomersUC,
		createCustomerUC,
		getCustomerUC,
		updateCustomerUC,
		deleteCustomerUC,
	)
	orderController := controller.NewOrderController(listOrdersUC, createOrderUC)

	// Handlers
	productHandler := handler.NewProductHandler(productController)
	customerHandler := handler.NewCustomerHandler(customerController)
	orderHandler := handler.NewOrderHandler(orderController)

	return &route.Handlers{
		Product:  productHandler,
		Customer: customerHandler,
		Order:    orderHandler,
	}
}
