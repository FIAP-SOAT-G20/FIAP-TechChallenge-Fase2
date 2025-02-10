// cmd/api/main.go
package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

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

//	@title			Clean Architecture API
//	@version		1.0
//	@description	API em Go usando Clean Architecture.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Seu Nome
//	@contact.url	http://seu-site.com
//	@contact.email	seu-email@email.com

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// Carrega as variáveis de ambiente
	envFile := getEnvFile()
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	// Carrega configurações
	cfg := config.LoadConfig(envFile)

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

func getEnvFile() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	filename := ".env"
	if env != "development" {
		filename = ".env." + env
	}

	// Try to find the env file in the current directory or parent directories
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for {
		path := filepath.Join(dir, filename)
		if _, err := os.Stat(path); err == nil {
			return path
		}

		// Get the parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// We've reached the root directory
			break
		}
		dir = parent
	}

	// If we haven't found the file, return the default path
	return ".env"
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
