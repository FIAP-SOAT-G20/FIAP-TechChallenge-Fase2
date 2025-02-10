.PHONY: all build run test clean help

# Variáveis
APP_NAME=app
MAIN_FILE=cmd/server/main.go

# Go comandos
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

# Ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make build        - Compila o projeto"
	@echo "  make run          - Executa o projeto"
	@echo "  make test         - Executa os testes"
	@echo "  make coverage     - Executa os testes com cobertura"
	@echo "  make clean        - Remove arquivos de build"
	@echo "  make mock         - Gera os mocks"
	@echo "  make lint         - Executa o linter"
	@echo "  make migrate-up   - Executa as migrações"
	@echo "  make migrate-down - Desfaz as migrações"
	@echo "  make install  	   - Instala as dependências"

# Build
build:
	$(GOBUILD) -o $(APP_NAME) $(MAIN_FILE)

# Run
run:
	$(GORUN) $(MAIN_FILE)

# Test
test:
	$(GOTEST) ./... -v

# Test com cobertura
coverage:
	$(GOTEST) ./... -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out

# Clean
clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)
	rm -f coverage.out

# Mockgen
mock:
	mockgen -source=internal/core/port/product_repository_port.go -destination=internal/core/port/mocks/product_repository_mock.go
	mockgen -source=internal/core/port/product_presenter_port.go -destination=internal/core/port/mocks/product_presenter_mock.go
	mockgen -source=internal/core/port/product_usecase_port.go -destination=internal/core/port/mocks/product_usecase_mock.go

# Lint
lint:
	golangci-lint run

# Migrate
migrate-up:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" up

migrate-down:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" down

install:
	go mod download
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install go.uber.org/mock/mockgen@latest
