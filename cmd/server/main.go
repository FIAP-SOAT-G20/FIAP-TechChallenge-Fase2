// cmd/api/main.go
package main

import (
	"log/slog"
	"os"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/datasources"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/gateway"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/product"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/database"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/routes"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/server"
)

//	@title			FIAP Tech Challenge Fase 2 - 10SOAT - G18
//	@version		1
//	@description	### API de um Fast Food para o Tech Challenge da FIAP - Fase 2 - 10SOAT - G18

//	@externalDocs.description	GitHub Repository
//	@externalDocs.url			https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	// Carrega configurações
	cfg := config.LoadConfig()

	// Inicializa o logger
	logger := setupLogger(cfg.Environment)

	// Inicializa o banco de dados
	db, err := database.NewPostgresConnection(cfg, logger)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	// Roda as migrações
	if err := db.Migrate(); err != nil {
		logger.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}

	// Inicializa as dependências e handlers
	handlers := setupHandlers(db)

	// Inicializa e inicia o servidor
	srv := server.NewServer(cfg, logger, handlers)
	if err := srv.Start(); err != nil {
		logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

func setupHandlers(db *database.Database) *routes.Handlers {
	// Datasources
	productDS := datasources.NewProductDataSource(db.DB)

	// Repositories
	productRepo := gateway.NewProductGateway(productDS)

	// Presenters
	productPresenter := presenter.NewProductPresenter()

	// Use cases
	listProductsUC := product.NewListProductsUseCase(productRepo, productPresenter)
	createProductUC := product.NewCreateProductUseCase(productRepo, productPresenter)
	getProductUC := product.NewGetProductUseCase(productRepo, productPresenter)
	updateProduct := product.NewUpdateProductUseCase(productRepo, productPresenter)
	deleteProduct := product.NewDeleteProductUseCase(productRepo)

	// Controllers
	productController := controller.NewProductController(
		listProductsUC,
		createProductUC,
		getProductUC,
		updateProduct,
		deleteProduct,
	)

	// Handlers
	productHandler := handler.NewProductHandler(productController)

	return &routes.Handlers{
		Product: productHandler,
	}
}

func setupLogger(environment string) *slog.Logger {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	if environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
