# Variáveis
APP_NAME=app
MAIN_FILE=cmd/server/main.go
DOCKER_REGISTRY=your-registry
VERSION=$(shell git describe --tags --always --dirty)
NAMESPACE=tech-challenge-system

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
	@echo "  make swagger      - Gera a documentação Swagger"
	@echo "  make lint         - Executa o linter"
	@echo "  make migrate-up   - Executa as migrações"
	@echo "  make migrate-down - Desfaz as migrações"
	@echo "  make install      - Instala as dependências"
	@echo "  make docker-build - Builda a imagem Docker"
	@echo "  make docker-push  - Publica a imagem no registry"
	@echo "  make k8s-apply    - Aplica manifestos Kubernetes"
	@echo "  make k8s-delete   - Remove recursos Kubernetes"
	@echo "  make k8s-logs     - Mostra logs da aplicação"
	@echo "  make k8s-status   - Mostra status dos recursos"

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
	mockgen -source=internal/core/port/product_gateway_port.go -destination=internal/core/port/mocks/product_gateway_mock.go
	mockgen -source=internal/core/port/product_presenter_port.go -destination=internal/core/port/mocks/product_presenter_mock.go
	mockgen -source=internal/core/port/product_usecase_port.go -destination=internal/core/port/mocks/product_usecase_mock.go

# Swagger
swagger:
	swag fmt ./...
	swag init -g ${MAIN_FILE} --parseInternal true

# Lint
lint:
	golangci-lint run

# Migrate
migrate-up:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" up

migrate-down:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" down

# Install
install:
	go mod download
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install go.uber.org/mock/mockgen@latest

# Docker
docker-build:
	docker build -t $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) .
	docker tag $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) $(DOCKER_REGISTRY)/$(APP_NAME):latest

docker-push:
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):latest

# Kubernetes
k8s-apply:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/config/
	kubectl apply -f k8s/postgres/
	kubectl apply -f k8s/app/

k8s-delete:
	kubectl delete -f k8s/app/
	kubectl delete -f k8s/postgres/
	kubectl delete -f k8s/config/
	kubectl delete -f k8s/namespace.yaml

k8s-logs:
	kubectl logs -f -l app=product-api -n $(NAMESPACE)

k8s-status:
	@echo "=== Pods ==="
	kubectl get pods -n $(NAMESPACE)
	@echo "\n=== Services ==="
	kubectl get svc -n $(NAMESPACE)
	@echo "\n=== Deployments ==="
	kubectl get deploy -n $(NAMESPACE)
	@echo "\n=== HPA ==="
	kubectl get hpa -n $(NAMESPACE)

# Dev environment
dev-up:
	docker-compose up -d

dev-down:
	docker-compose down

# Security
scan:
	govulncheck ./...
	trivy image $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)

# All
all: lint test build docker-build