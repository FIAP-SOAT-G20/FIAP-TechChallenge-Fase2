# <p align="center">FIAP Tech Challenge 2 - G20 Fast Food</p>

<p align="center">
    <img src="https://img.shields.io/badge/Code-Go-informational?style=flat-square&logo=go&color=00ADD8" alt="Go" />
    <img src="https://img.shields.io/badge/Tools-Gin-informational?style=flat-square&logo=go&color=00ADD8" alt="Gin" />
    <img src="https://img.shields.io/badge/Tools-PostgreSQL-informational?style=flat-square&logo=postgresql&color=4169E1" alt="PostgreSQL" />
    <img src="https://img.shields.io/badge/Tools-Swagger-informational?style=flat-square&logo=swagger&color=85EA2D" alt="Swagger" />
    <img src="https://img.shields.io/badge/Tools-Docker-informational?style=flat-square&logo=docker&color=2496ED" alt="Docker" />
    <img src="https://img.shields.io/badge/Tools-Kubernetes-informational?style=flat-square&logo=kubernetes&color=326CE5" alt="Kubernetes" />
</p>

## 🏗️ Arquitetura
### **1️⃣ Core (Camada mais interna)**
- `domain/`: Entidades e regras de negócio centrais.
- `usecases/`: Casos de uso da aplicação.
- `ports/`: Interfaces que definem contratos entre camadas, garantindo independência.

### **2️⃣ Adapters (Camada intermediária)**
- `controller/`: Coordena o fluxo de dados entre use cases e infraestrutura.
- `presenter/`: Formata dados para apresentação.
- `gateway/`: Implementa acesso a dados de fontes externas (banco de dados, APIs, etc.).
- `datasources/`: Implementações concretas de fontes de dados.

### **3️⃣ Infrastructure (Camada externa)**
- `config/`: Gerenciamento de configurações da aplicação.
- `database/`: Configuração e conexão com o banco de dados.
- `server/`: Inicialização do servidor HTTP.
- `routes/`: Definição das rotas da API.
- `middleware/`: Middlewares HTTP para tratamento de requisições.
- `logger/`: Logger estruturado para logs detalhados.
- `handler/`: Tratamento das requisições HTTP.

## 🛠️ Tecnologias Utilizadas

- [Go 1.23+](https://golang.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [slog](https://pkg.go.dev/log/slog) para logs estruturados
- [godotenv](https://github.com/joho/godotenv) para gerenciamento de variáveis de ambiente
- [go-playground/validator](https://github.com/go-playground/validator) para validações estruturadas
- [Docker](https://www.docker.com/) para containerização
- [Kubernetes](https://kubernetes.io/) para orquestração
- [Swagger](https://swagger.io/) para documentação da API

## 📋 Pré-requisitos

- Go 1.23 ou superior
- PostgreSQL
- Docker e Docker Compose (opcional, mas recomendado)
- Make (opcional, para automação de comandos)

## 🚀 Configuração e Execução

1. Clone o repositório:
```bash
git clone https://github.com/your-username/your-project.git
cd your-project
```

2. Copie o arquivo de exemplo de variáveis de ambiente:
```bash
cp .env.example .env
```

3. Configure as variáveis de ambiente no arquivo `.env`.

4. Instale as dependências:
```bash
go mod download
```

5. Execute o projeto:
```bash
go run cmd/server/main.go
```

## 🛠️ Docker e Kubernetes

### **Executando com Docker Compose**

```bash
docker-compose up --build
```

### **Executando com Kubernetes**

```bash
kubectl apply -f k8s/
```

## 📝 Documentação da API

A documentação da API está disponível via Swagger:
- URL local: `http://localhost:8080/docs/index.html`

## 🏗️ Makefile

O projeto inclui um `Makefile` para facilitar a execução de comandos comuns. Os principais comandos disponíveis são:

```bash
make run              # Executa a aplicação
make build            # Compila o projeto
make test             # Executa os testes
make lint             # Verifica a qualidade do código
make docker-build     # Constrói a imagem Docker
make docker-run       # Executa o container Docker
make k8s-deploy       # Implanta a aplicação no Kubernetes
```

Para executar um comando, basta rodar:
```bash
make <comando>
```

## 🧪 Testes

Execute os testes unitários:
```bash
go test ./...
```

Para ver a cobertura de testes:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```