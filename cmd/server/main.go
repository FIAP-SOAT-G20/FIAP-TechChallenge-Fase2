package main

import (
	"os"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/gateway"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/payment"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/database"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/datasource"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/logger"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/route"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/server"
)

// @title						FIAP Tech Challenge Fase 2 - 10SOAT - G18
// @version					1
// @description				### API de um Fast Food para o Tech Challenge da FIAP - Fase 2
// @servers					[ { "url": "http://localhost:8080" } ]
// @host						localhost:8080
// @BasePath					/api/v1
// @tag.name					sign-up
// @tag.description			Regiter a new customer or staff
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
	// Datasources
	productDS := datasource.NewProductDataSource(db.DB)
	customerDS := datasource.NewCustomerDataSource(db.DB)
	orderDS := datasource.NewOrderDataSource(db.DB)
	orderProductDS := datasource.NewOrderProductDataSource(db.DB)
	staffDS := datasource.NewStaffDataSource(db.DB)
	paymentDS := datasource.NewPaymentDataSource(db.DB)

	// Gateways
	productGateway := gateway.NewProductGateway(productDS)
	customerGateway := gateway.NewCustomerGateway(customerDS)
	orderGateway := gateway.NewOrderGateway(orderDS)
	orderProductGateway := gateway.NewOrderProductGateway(orderProductDS)
	staffGateway := gateway.NewStaffGateway(staffDS)

	// Use cases
	productUC := usecase.NewProductUseCase(productGateway)
	customerUC := usecase.NewCustomerUseCase(customerGateway)
	orderUC := usecase.NewOrderUseCase(orderGateway)
	orderProductUC := usecase.NewOrderProductUseCase(orderProductGateway)
	staffUC := usecase.NewStaffUseCase(staffGateway)

	// Controllers
	productController := controller.NewProductController(productUC)
	customerController := controller.NewCustomerController(customerUC)
	orderController := controller.NewOrderController(orderUC)
	orderProductController := controller.NewOrderProductController(orderProductUC)
	staffController := controller.NewStaffController(staffUC)
	paymentGateway := gateway.NewPaymentGateway(paymentDS)

	// Presenters
	productPresenter := presenter.NewProductJsonPresenter()
	// productPresenter := presenter.NewProductXmlPresenter()
	customerPresenter := presenter.NewCustomerJsonPresenter()
	orderPresenter := presenter.NewOrderJsonPresenter()
	orderProductPresenter := presenter.NewOrderProductJsonPresenter()
	paymentPresenter := presenter.NewPaymentJsonPresenter()

	// Externals
	paymentExternal := datasource.NewPaymentExternal()

	// Use cases - Product
	listProductsUC := product.NewListProductsUseCase(productGateway, productPresenter)
	createProductUC := product.NewCreateProductUseCase(productGateway, productPresenter)
	getProductUC := product.NewGetProductUseCase(productGateway, productPresenter)
	updateProductUC := product.NewUpdateProductUseCase(productGateway, productPresenter)
	deleteProductUC := product.NewDeleteProductUseCase(productGateway, productPresenter)
	// Use cases - Customer
	listCustomersUC := customer.NewListCustomersUseCase(customerGateway, customerPresenter)
	createCustomerUC := customer.NewCreateCustomerUseCase(customerGateway, customerPresenter)
	getCustomerUC := customer.NewGetCustomerUseCase(customerGateway, customerPresenter)
	updateCustomerUC := customer.NewUpdateCustomerUseCase(customerGateway, customerPresenter)
	deleteCustomerUC := customer.NewDeleteCustomerUseCase(customerGateway, customerPresenter)
	// Use cases - Order
	listOrdersUC := order.NewListOrdersUseCase(orderGateway, orderPresenter)
	createOrderUC := order.NewCreateOrderUseCase(orderGateway, orderPresenter)
	getOrderUC := order.NewGetOrderUseCase(orderGateway, orderPresenter)
	updateOrderUC := order.NewUpdateOrderUseCase(orderGateway, orderPresenter)
	deleteOrderUC := order.NewDeleteOrderUseCase(orderGateway, orderPresenter)
	// Use cases - OrderProduct
	listOrderProductsUC := orderproduct.NewListOrderProductsUseCase(orderProductGateway, orderProductPresenter)
	createOrderProductUC := orderproduct.NewCreateOrderProductUseCase(orderProductGateway, orderProductPresenter)
	getOrderProductUC := orderproduct.NewGetOrderProductUseCase(orderProductGateway, orderProductPresenter)
	updateOrderProductUC := orderproduct.NewUpdateOrderProductUseCase(orderProductGateway, orderProductPresenter)
	deleteOrderProductUC := orderproduct.NewDeleteOrderProductUseCase(orderProductGateway, orderProductPresenter)
	// Use cases - Payment
	createPaymentUC := payment.NewCreatePaymentUseCase(paymentGateway, orderGateway, paymentExternal, paymentPresenter)

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
	orderController := controller.NewOrderController(
		listOrdersUC,
		createOrderUC,
		getOrderUC,
		updateOrderUC,
		deleteOrderUC,
	)
	orderProductController := controller.NewOrderProductController(
		listOrderProductsUC,
		createOrderProductUC,
		getOrderProductUC,
		updateOrderProductUC,
		deleteOrderProductUC,
	)
	paymentController := controller.NewPaymentController(
		createPaymentUC,
	)

	// Handlers
	productHandler := handler.NewProductHandler(productController)
	customerHandler := handler.NewCustomerHandler(customerController)
	orderHandler := handler.NewOrderHandler(orderController)
	orderProductHandler := handler.NewOrderProductHandler(orderProductController)
	staffHandler := handler.NewStaffHandler(staffController)
	healthCheckHandler := handler.NewHealthCheckHandler()
	paymentHandler := handler.NewPaymentHandler(paymentController)

	return &route.Handlers{
		Product:      productHandler,
		Customer:     customerHandler,
		Order:        orderHandler,
		OrderProduct: orderProductHandler,
		Staff:        staffHandler,
		HealthCheck:  healthCheckHandler,
		Payment:      paymentHandler,
	}
}
