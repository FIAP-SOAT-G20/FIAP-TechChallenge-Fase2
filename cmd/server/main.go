package main

import (
	"os"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/staff"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/gateway"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/customer"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order"
	orderproduct "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/order_product"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/product"
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
	// Datasource's
	productDS := datasource.NewProductDataSource(db.DB)
	customerDS := datasource.NewCustomerDataSource(db.DB)
	staffDS := datasource.NewStaffDataSource(db.DB)
	orderDS := datasource.NewOrderDataSource(db.DB)
	orderProductDS := datasource.NewOrderProductDataSource(db.DB)

	// Gateways
	productGateway := gateway.NewProductGateway(productDS)
	customerGateway := gateway.NewCustomerGateway(customerDS)
	staffGateway := gateway.NewStaffGateway(staffDS)
	orderGateway := gateway.NewOrderGateway(orderDS)
	orderProductGateway := gateway.NewOrderProductGateway(orderProductDS)

	// Use cases
	productUC := product.NewProductUseCase(productGateway)
	staffUC := staff.NewStaffUseCase(staffGateway)
	customerUC := customer.NewCustomerUseCase(customerGateway)
	// Use cases - Order
	listOrdersUC := order.NewListOrdersUseCase(orderGateway)
	createOrderUC := order.NewCreateOrderUseCase(orderGateway)
	getOrderUC := order.NewGetOrderUseCase(orderGateway)
	updateOrderUC := order.NewUpdateOrderUseCase(orderGateway)
	deleteOrderUC := order.NewDeleteOrderUseCase(orderGateway)
	// Use cases - OrderProduct
	listOrderProductsUC := orderproduct.NewListOrderProductsUseCase(orderProductGateway)
	createOrderProductUC := orderproduct.NewCreateOrderProductUseCase(orderProductGateway)
	getOrderProductUC := orderproduct.NewGetOrderProductUseCase(orderProductGateway)
	updateOrderProductUC := orderproduct.NewUpdateOrderProductUseCase(orderProductGateway)
	deleteOrderProductUC := orderproduct.NewDeleteOrderProductUseCase(orderProductGateway)

	// Controllers
	productController := controller.NewProductController(productUC)
	customerController := controller.NewCustomerController(customerUC)
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
	staffController := controller.NewStaffController(staffUC)

	// Handlers
	productHandler := handler.NewProductHandler(productController)
	customerHandler := handler.NewCustomerHandler(customerController)

	staffHandler := handler.NewStaffHandler(staffController)
	orderHandler := handler.NewOrderHandler(orderController)
	orderProductHandler := handler.NewOrderProductHandler(orderProductController)

	return &route.Handlers{
		Product:      productHandler,
		Customer:     customerHandler,
		Staff:        staffHandler,
		Order:        orderHandler,
		OrderProduct: orderProductHandler,
	}
}
