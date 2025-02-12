package main

import (
	"os"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/gateway"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/product"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/database"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/datasources"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/logger"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/routes"
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

func setupHandlers(db *database.Database) *routes.Handlers {
	// Datasource's
	productDS := datasources.NewProductDataSource(db.DB)

	// Gateways
	productGateway := gateway.NewProductGateway(productDS)

	// Presenters
	productPresenter := presenter.NewProductJsonPresenter()
	// productPresenter := presenter.NewProductXmlPresenter()

	// Use cases
	listProductsUC := product.NewListProductsUseCase(productGateway, productPresenter)
	createProductUC := product.NewCreateProductUseCase(productGateway, productPresenter)
	getProductUC := product.NewGetProductUseCase(productGateway, productPresenter)
	updateProductUC := product.NewUpdateProductUseCase(productGateway, productPresenter)
	deleteProductUC := product.NewDeleteProductUseCase(productGateway, productPresenter)

	// Controllers
	productController := controller.NewProductController(
		listProductsUC,
		createProductUC,
		getProductUC,
		updateProductUC,
		deleteProductUC,
	)

	// Handlers
	productHandler := handler.NewProductHandler(productController)

	return &routes.Handlers{
		Product: productHandler,
	}
}
