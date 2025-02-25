.DEFAULT_GOAL := help

APP_NAME=app
MAIN_FILE=cmd/server/main.go
DOCKER_REGISTRY=your-registry
VERSION=$(shell git describe --tags --always --dirty)
NAMESPACE=tech-challenge-system
TEST_PATH=./internal/...
TEST_COVERAGE_FILE_NAME=coverage.out
MIGRATION_PATH = internal/infrastructure/database/migrations
DB_URL = postgres://postgres:postgres@localhost:5432/fastfood_10soat_g18_tc2?sslmode=disable

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

# Looks at comments using ## on targets and uses them to produce a help output.
.PHONY: help
help: ALIGN=22
help: ## Print this message
	@echo "Usage: make <command>"
	@awk -F '::? .*## ' -- "/^[^':]+::? .*## /"' { printf "  make '$$(tput bold)'%-$(ALIGN)s'$$(tput sgr0)' - %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build the application
	@echo  "🟢 Building the application..."
	$(GOCMD) fmt ./...
	$(GOBUILD) -o bin/$(APP_NAME) $(MAIN_FILE)

.PHONY: run-db
run-db: ## Run the database
	@echo  "🟢 Running the database..."
	docker-compose up -d db dbadmin

.PHONY: run
run: build run-db ## Run the application
	@echo  "🟢 Running the application..."
	$(GORUN) $(MAIN_FILE) || true

.PHONY: stop
stop: ## Stop the application
	@echo  "🔴 Stopping the application..."
	docker-compose down	

.PHONY: stop-db
stop-db: ## Stop the database
	@echo  "🔴 Stopping the database..."
	docker-compose down db dbadmin

.PHONY: run-air
run-air: build ## Run the application with Air
	@echo  "🟢 Running the application with Air..."
	air -c air.toml

.PHONY: test
test: ## Run tests
	@echo  "🟢 Running tests..."
	$(GOTEST) $(TEST_PATH) -v

.PHONY: coverage
coverage: ## Run tests with coverage
	@echo  "🟢 Running tests with coverage..."
	$(GOTEST) $(TEST_PATH) -coverprofile=$(TEST_COVERAGE_FILE_NAME).tmp
	@cat $(TEST_COVERAGE_FILE_NAME).tmp | grep -v "_mock.go" > $(TEST_COVERAGE_FILE_NAME)
	@rm $(TEST_COVERAGE_FILE_NAME).tmp
	$(GOCMD) tool cover -html=$(TEST_COVERAGE_FILE_NAME)

.PHONY: clean
clean: ## Clean up binaries and coverage files
	@echo "🔴 Cleaning up..."
	$(GOCLEAN)
	rm -f $(APP_NAME)
	rm -f coverage.out

.PHONY: mock
mock: ## Generate mocks
	@echo  "🟢 Generating mocks..."
	mockgen -source=internal/core/port/presenter_port.go -destination=internal/core/port/mocks/presenter_mock.go
	mockgen -source=internal/core/port/product_gateway_port.go -destination=internal/core/port/mocks/product_gateway_mock.go
	mockgen -source=internal/core/port/product_usecase_port.go -destination=internal/core/port/mocks/product_usecase_mock.go
	mockgen -source=internal/core/port/product_controller_port.go -destination=internal/core/port/mocks/product_controller_mock.go
	mockgen -source=internal/core/port/customer_gateway_port.go -destination=internal/core/port/mocks/customer_gateway_mock.go
	mockgen -source=internal/core/port/customer_usecase_port.go -destination=internal/core/port/mocks/customer_usecase_mock.go
	mockgen -source=internal/core/port/customer_controller_port.go -destination=internal/core/port/mocks/customer_controller_mock.go
	mockgen -source=internal/core/port/order_gateway_port.go -destination=internal/core/port/mocks/order_gateway_mock.go
	mockgen -source=internal/core/port/order_usecase_port.go -destination=internal/core/port/mocks/order_usecase_mock.go
	mockgen -source=internal/core/port/order_controller_port.go -destination=internal/core/port/mocks/order_controller_mock.go
	mockgen -source=internal/core/port/order_product_gateway_port.go -destination=internal/core/port/mocks/order_product_gateway_mock.go
	mockgen -source=internal/core/port/order_product_controller_port.go -destination=internal/core/port/mocks/order_product_controller_mock.go
	mockgen -source=internal/core/port/staff_gateway_port.go -destination=internal/core/port/mocks/staff_gateway_mock.go
	mockgen -source=internal/core/port/staff_usecase_port.go -destination=internal/core/port/mocks/staff_usecase_mock.go
	mockgen -source=internal/core/port/staff_controller_port.go -destination=internal/core/port/mocks/staff_controller_mock.go
	mockgen -source=internal/core/port/order_history_gateway_port.go -destination=internal/core/port/mocks/order_history_gateway_mock.go
	mockgen -source=internal/core/port/order_history_usecase_port.go -destination=internal/core/port/mocks/order_history_usecase_mock.go
	mockgen -source=internal/core/port/order_history_controller_port.go -destination=internal/core/port/mocks/order_history_controller_mock.go
	mockgen -source=internal/core/port/payment_gateway_port.go -destination=internal/core/port/mocks/payment_gateway_mock.go
	mockgen -source=internal/core/port/payment_usecase_port.go -destination=internal/core/port/mocks/payment_usecase_mock.go
	mockgen -source=internal/core/port/payment_controller_port.go -destination=internal/core/port/mocks/payment_controller_mock.go
	

.PHONY: swagger
swagger: ## Generate Swagger documentation
	@echo  "🟢 Generating Swagger documentation..."
	swag fmt ./...
	swag init -g ${MAIN_FILE} --parseInternal true

.PHONY: lint
lint: ## Run linter
	@echo  "🟢 Running linter..."
	golangci-lint run

.PHONY: migrate-create
migrate-create: ## Create new migration, usage example: make migrate-create name=create_table_products
	@echo  "🟢 Creating new migration..."
# if name is not passed, required argument
ifndef name
	$(error name is not set, usage example: make migrate-create name=create_table_products)
endif
	migrate create -ext sql -dir ${MIGRATION_PATH} -seq $(name)

.PHONY: migrate-up
migrate-up: ## Run migrations
	@echo  "🟢 Running migrations..."
	migrate -path ${MIGRATION_PATH} -database "${DB_URL}" -verbose up

.PHONY: migrate-down
migrate-down: ## Roll back migrations
	@echo  "🔴 Rolling back migrations..."
	migrate -path ${MIGRATION_PATH} -database "${DB_URL}" -verbose down

.PHONY: install
install: ## Install dependencies
	@echo  "🟢 Installing dependencies..."
	go mod download
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install go.uber.org/mock/mockgen@latest
	@go install github.com/air-verse/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: docker-build
docker-build: ## Build Docker image
	@echo  "🟢 Building Docker image..."
	docker build --platform linux/amd64 -t $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) .
	docker tag $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) $(DOCKER_REGISTRY)/$(APP_NAME):latest

.PHONY: docker-push
docker-push: ## Push Docker image
	@echo  "🟢 Pushing Docker image..."
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):latest

.PHONY: k8s-apply
k8s-apply: ## Apply Kubernetes manifests
	@echo  "🟢 Applying Kubernetes manifests..."
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/config/
	kubectl apply -f k8s/postgres/
	kubectl apply -f k8s/app/

.PHONY: k8s-delete
k8s-delete: ## Delete Kubernetes resources
	@echo  "🔴 Deleting Kubernetes resources..."
	kubectl delete -f k8s/app/
	kubectl delete -f k8s/postgres/
	kubectl delete -f k8s/config/
	kubectl delete -f k8s/namespace.yaml

.PHONY: k8s-logs
k8s-logs: ## Show application logs
	@echo  "🟢 Showing application logs..."
	kubectl logs -f -l app=product-api -n $(NAMESPACE)

.PHONY: k8s-status
k8s-status: ## Show Kubernetes resources status
	@echo  "🟢 Showing Kubernetes resources status..."
	@echo "=== Pods ==="
	kubectl get pods -n $(NAMESPACE)
	@echo "\n=== Services ==="
	kubectl get svc -n $(NAMESPACE)
	@echo "\n=== Deployments ==="
	kubectl get deploy -n $(NAMESPACE)
	@echo "\n=== HPA ==="
	kubectl get hpa -n $(NAMESPACE)

.PHONY: compose-build
compose-build: ## Build the application with Docker Compose
	@echo "🟢 Building the application..."
	docker compose build

.PHONY: compose-up
compose-up: ## Start development environment with Docker Compose
	@echo  "🟢 Starting development environment..."
	docker-compose up -d --wait --build

.PHONY: compose-down
compose-down: ## Stop development environment with Docker Compose
	@echo  "🔴 Stopping development environment..."
	docker-compose down

.PHONY: compose-clean
compose-clean: ## Clean the application with Docker Compose, removing volumes and images
	@echo "🔴 Cleaning the application..."
	docker compose down --volumes --rmi all

.PHONY: scan
scan: ## Run security scan
	@echo  "🟢 Running security scan..."
	govulncheck -show verbose ./...
#	trivy image $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) # TODO: Enable when the image is available

.PHONY: new-branch
new-branch: ## Create new branch
	@echo "🟢 Creating new branch..."
	./scripts/new-branch.sh -c

.PHONY: pull-request
pull-request: ## Create pull request
	@echo "🟢 Creating pull request..."
	./scripts/pull-request.sh
