# Build stage
FROM golang:1.23-alpine AS builder

# Configuração do diretório de trabalho
WORKDIR /app

# Copia go.mod e go.sum
COPY go.mod go.sum ./

# Download das dependências
RUN go mod download

# Copia o código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copia o binário do builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expõe a porta da aplicação
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]