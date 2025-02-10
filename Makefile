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
	@echo "Usage: make <command>"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make run-air      - Run the application with Air"
	@echo "  make test         - Run tests"
	@echo "  make coverage     - Run tests with coverage"
	@echo "  make clean        - Clean up"
	@echo "  make mock         - Generate mocks"
	@echo "  make swagger      - Generate Swagger documentation"
	@echo "  make lint         - Run linter"
	@echo "  make migrate-up   - Run migrations"
	@echo "  make migrate-down - Roll back migrations"
	@echo "  make install      - Install dependencies"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-push  - Push Docker image"
	@echo "  make k8s-apply    - Apply Kubernetes manifests"
	@echo "  make k8s-delete   - Delete Kubernetes resources"
	@echo "  make k8s-logs     - Show application logs"
	@echo "  make k8s-status   - Show Kubernetes resources status"

build:
	@echo  "🟢 Building the application..."
	$(GOBUILD) fmt ./...
	$(GOBUILD) -o $(APP_NAME) $(MAIN_FILE)

run:
	@echo  "🟢 Running the application..."
	docker-compose up -d postgres
	$(GORUN) $(MAIN_FILE)

run-air:
	@echo  "🟢 Running the application with Air..."
	air -c air.toml

test:
	@echo  "🟢 Running tests..."
	$(GOTEST) ./... -v

coverage:
	@echo  "🟢 Running tests with coverage..."
	$(GOTEST) ./... -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out

clean:
	@echo "🔴 Cleaning up..."
	$(GOCLEAN)
	rm -f $(APP_NAME)
	rm -f coverage.out

mock:
	@echo  "🟢 Generating mocks..."
	mockgen -source=internal/core/port/product_gateway_port.go -destination=internal/core/port/mocks/product_gateway_mock.go
	mockgen -source=internal/core/port/product_presenter_port.go -destination=internal/core/port/mocks/product_presenter_mock.go
	mockgen -source=internal/core/port/product_usecase_port.go -destination=internal/core/port/mocks/product_usecase_mock.go

swagger:
	@echo  "🟢 Generating Swagger documentation..."
	swag fmt ./...
	swag init -g ${MAIN_FILE} --parseInternal true

lint:
	@echo  "🟢 Running linter..."
	golangci-lint run

migrate-up:
	@echo  "🟢 Running migrations..."
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" up

migrate-down:
	@echo  "🔴 Rolling back migrations..."
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" down

install:
	@echo  "🟢 Installing dependencies..."
	go mod download
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install go.uber.org/mock/mockgen@latest

docker-build:
	@echo  "🟢 Building Docker image..."
	docker build -t $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) .
	docker tag $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) $(DOCKER_REGISTRY)/$(APP_NAME):latest

docker-push:
	@echo  "🟢 Pushing Docker image..."
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):latest

k8s-apply:
	@echo  "🟢 Applying Kubernetes manifests..."
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/config/
	kubectl apply -f k8s/postgres/
	kubectl apply -f k8s/app/

k8s-delete:
	@echo  "🔴 Deleting Kubernetes resources..."
	kubectl delete -f k8s/app/
	kubectl delete -f k8s/postgres/
	kubectl delete -f k8s/config/
	kubectl delete -f k8s/namespace.yaml

k8s-logs:
	@echo  "🟢 Showing application logs..."
	kubectl logs -f -l app=product-api -n $(NAMESPACE)

k8s-status:
	@echo  "🟢 Showing Kubernetes resources status..."
	@echo "=== Pods ==="
	kubectl get pods -n $(NAMESPACE)
	@echo "\n=== Services ==="
	kubectl get svc -n $(NAMESPACE)
	@echo "\n=== Deployments ==="
	kubectl get deploy -n $(NAMESPACE)
	@echo "\n=== HPA ==="
	kubectl get hpa -n $(NAMESPACE)

dev-build:
	@echo "🟢 Building the application with docker compose..."
	docker compose build

dev-up:
	@echo  "🟢 Starting development environment..."
	docker-compose up -d --wait

dev-down:
	@echo  "🔴 Stopping development environment..."
	docker-compose down

dev-clean:
	echo "🔴 Cleaning the application ..."
	docker compose down --volumes --rmi all

scan:
	@echo  "🟢 Running security scan..."
	govulncheck ./...
	trivy image $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
