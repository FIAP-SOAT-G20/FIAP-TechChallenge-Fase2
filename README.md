# Clean Architecture em Go

Este projeto é uma implementação de uma API REST seguindo os princípios da Clean Architecture em Go. O projeto utiliza tecnologias modernas e boas práticas de desenvolvimento.

## 🏗️ Arquitetura

O projeto segue a Clean Architecture, dividida em camadas:

### Core (Camada mais interna)
- `domain/`: Entidades e regras de negócio centrais
- `usecases/`: Casos de uso da aplicação
- `ports/`: Interfaces que definem contratos entre camadas

### Adapters (Camada intermediária)
- `controllers/`: Coordena o fluxo de dados
- `presenters/`: Formata dados para apresentação
- `repositories/`: Implementa acesso a dados
- `handlers/`: Lida com requisições HTTP

### Infrastructure (Camada externa)
- `config/`: Configurações da aplicação
- `database/`: Conexão com banco de dados
- `server/`: Servidor HTTP
- `routes/`: Definição de rotas
- `middleware/`: Middlewares HTTP

## 🛠️ Tecnologias

- [Go 1.23+](https://golang.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [slog](https://pkg.go.dev/log/slog) para logs estruturados
- [godotenv](https://github.com/joho/godotenv) para variáveis de ambiente

## 📋 Pré-requisitos

- Go 1.21 ou superior
- PostgreSQL
- Make (opcional, para usar os comandos do Makefile)

## 🚀 Configuração e Execução

1. Clone o repositório
```bash
git clone https://github.com/your-username/your-project.git
cd your-project
```

2. Copie o arquivo de exemplo de variáveis de ambiente
```bash
cp .env.example .env
```

3. Configure as variáveis de ambiente no arquivo `.env`:
```env
# Database settings
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=products
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=25
DB_CONN_MAX_LIFETIME=5m

# Server settings
SERVER_PORT=8080
SERVER_READ_TIMEOUT=10s
SERVER_WRITE_TIMEOUT=10s
SERVER_IDLE_TIMEOUT=60s

# Environment
ENVIRONMENT=development
```

4. Instale as dependências
```bash
go mod download
```

5. Execute o projeto
```bash
go run cmd/api/main.go
```

## 📝 Documentação da API

### Produtos

#### Listar Produtos
```http
GET /api/v1/products
```

Parâmetros de Query:
- `page` (opcional): Número da página (default: 1)
- `limit` (opcional): Itens por página (default: 10)
- `name` (opcional): Filtrar por nome
- `category_id` (opcional): Filtrar por categoria

#### Criar Produto
```http
POST /api/v1/products
```

Body:
```json
{
  "name": "Produto Exemplo",
  "description": "Descrição do produto",
  "price": 99.99,
  "category_id": 1
}
```

## 🧪 Testes

Execute os testes:
```bash
go test ./...
```

Para ver a cobertura de testes:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 📁 Estrutura de Diretórios

```
.
├── cmd/
│   └── api/              # Ponto de entrada da aplicação
├── internal/
│   ├── core/
│   │   ├── domain/      # Entidades e regras de negócio
│   │   ├── usecases/    # Casos de uso
│   │   └── ports/       # Interfaces/Contratos
│   ├── adapters/
│   │   ├── controllers/ # Controladores
│   │   ├── handlers/    # Handlers HTTP
│   │   ├── presenters/  # Formatadores de resposta
│   │   └── repositories/# Repositórios
│   └── infrastructure/
│       ├── config/      # Configurações
│       ├── database/    # Conexão com banco
│       ├── server/      # Servidor HTTP
│       └── middleware/  # Middlewares
└── scripts/             # Scripts úteis
```