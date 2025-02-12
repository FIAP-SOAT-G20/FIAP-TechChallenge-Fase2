APP_NAME=app
MAIN_FILE=cmd/server/main.go
DOCKER_REGISTRY=your-registry
VERSION=$(shell git describe --tags --always --dirty)
NAMESPACE=tech-challenge-system
TEST_PATH=./internal/...
TEST_COVERAGE_FILE_NAME=coverage.out

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

help:
	@echo "Usage: make <command>"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make run-air       - Run the application with Air"
	@echo "  make test          - Run tests"
	@echo "  make coverage      - Run tests with coverage"
	@echo "  make clean         - Clean up"
	@echo "  make mock          - Generate mocks"
	@echo "  make swagger       - Generate Swagger documentation"
	@echo "  make lint          - Run linter"
	@echo "  make migrate-up    - Run migrations"
	@echo "  make migrate-down  - Roll back migrations"
	@echo "  make install       - Install dependencies"
	@echo "  make scan          - Run security scan"
	@echo "  make new-branch    - Create new branch"
	@echo "  make pull-request  - Create pull request"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-push   - Push Docker image"
	@echo "  make compose-build - Build the application with Docker Compose"
	@echo "  make compose-up    - Start development environment with Docker Compose"
	@echo "  make compose-down  - Stop development environment with Docker Compose"
	@echo "  make compose-clean - Clean the application with Docker Compose"
	@echo "  make k8s-apply     - Apply Kubernetes manifests"
	@echo "  make k8s-delete    - Delete Kubernetes resources"
	@echo "  make k8s-logs      - Show application logs"
	@echo "  make k8s-status    - Show Kubernetes resources status"

.PHONY: build
build:
	@echo  "🟢 Building the application..."
	$(GOBUILD) fmt ./...
	$(GOBUILD) -o $(APP_NAME) $(MAIN_FILE)

.PHONY: run
run:
	@echo  "🟢 Running the application..."
	docker-compose up -d postgres
	$(GORUN) $(MAIN_FILE)

.PHONY: run-air
run-air:
	@echo  "🟢 Running the application with Air..."
	air -c air.toml

.PHONY: test
test:
	@echo  "🟢 Running tests..."
	$(GOTEST) $(TEST_PATH) -v

.PHONY: coverage
coverage:
	@echo  "🟢 Running tests with coverage..."
	$(GOTEST) $(TEST_PATH) -coverprofile=$(TEST_COVERAGE_FILE_NAME).tmp
	@cat $(TEST_COVERAGE_FILE_NAME).tmp | grep -v "_mock.go" > $(TEST_COVERAGE_FILE_NAME)
	@rm $(TEST_COVERAGE_FILE_NAME).tmp
	$(GOCMD) tool cover -html=$(TEST_COVERAGE_FILE_NAME)

.PHONY: clean
clean:
	@echo "🔴 Cleaning up..."
	$(GOCLEAN)
	rm -f $(APP_NAME)
	rm -f coverage.out

.PHONY: mock
mock:
	@echo  "🟢 Generating mocks..."
	mockgen -source=internal/core/port/product_gateway_port.go -destination=internal/core/port/mocks/product_gateway_mock.go
	mockgen -source=internal/core/port/product_presenter_port.go -destination=internal/core/port/mocks/product_presenter_mock.go
	mockgen -source=internal/core/port/product_usecase_port.go -destination=internal/core/port/mocks/product_usecase_mock.go
	mockgen -source=internal/adapter/dto/response_writer.go -destination=internal/adapter/dto/mocks/response_writer_mock.go

.PHONY: swagger
swagger:
	@echo  "🟢 Generating Swagger documentation..."
	swag fmt ./...
	swag init -g ${MAIN_FILE} --parseInternal true

.PHONY: lint
lint:
	@echo  "🟢 Running linter..."
	golangci-lint run

.PHONY: migrate-up
migrate-up:
	@echo  "🟢 Running migrations..."
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	@echo  "🔴 Rolling back migrations..."
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/products?sslmode=disable" down

.PHONY: install
install:
	@echo  "🟢 Installing dependencies..."
	go mod download
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install go.uber.org/mock/mockgen@latest

.PHONY: docker-build
docker-build:
	@echo  "🟢 Building Docker image..."
	docker build -t $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) .
	docker tag $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) $(DOCKER_REGISTRY)/$(APP_NAME):latest

.PHONY: docker-push
docker-push:
	@echo  "🟢 Pushing Docker image..."
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(APP_NAME):latest

.PHONY: k8s-apply
k8s-apply:
	@echo  "🟢 Applying Kubernetes manifests..."
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/config/
	kubectl apply -f k8s/postgres/
	kubectl apply -f k8s/app/

.PHONY: k8s-delete
k8s-delete:
	@echo  "🔴 Deleting Kubernetes resources..."
	kubectl delete -f k8s/app/
	kubectl delete -f k8s/postgres/
	kubectl delete -f k8s/config/
	kubectl delete -f k8s/namespace.yaml

.PHONY: k8s-logs
k8s-logs:
	@echo  "🟢 Showing application logs..."
	kubectl logs -f -l app=product-api -n $(NAMESPACE)

.PHONY: k8s-status
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

.PHONY: compose-build
compose-build:
	@echo "🟢 Building the application..."
	docker compose build

.PHONY: compose-up
compose-up:
	@echo  "🟢 Starting development environment..."
	docker-compose up -d --wait

.PHONY: compose-down
compose-down:
	@echo  "🔴 Stopping development environment..."
	docker-compose down

.PHONY: compose-clean
compose-clean:
	@echo "🔴 Cleaning the application..."
	docker compose down --volumes --rmi all

.PHONY: scan
scan:
	@echo  "🟢 Running security scan..."
	govulncheck -show verbose ./...
#	trivy image $(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION) # TODO: Enable when the image is available

.PHONY: new-branch
new-branch:
	@echo "🟢 Creating new branch..."
	./scripts/new-branch.sh -c

.PHONY: pull-request
pull-request:
	@echo "🟢 Creating pull request..."
	./scripts/pull-request.sh
