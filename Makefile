.PHONY: all build run test clean swagger

# Variáveis
APP_NAME=app
MAIN_FILE=cmd/server/main.go

# Go comandos
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

# Build
build:
	$(GOBUILD) -o $(APP_NAME) $(MAIN_FILE)

# Run
run:
	$(GORUN) $(MAIN_FILE)

# Test
test:
	$(GOTEST) ./... -v

# Clean
clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)

# Swagger
swagger:
	swag init -g $(MAIN_FILE) -o docs