FROM golang:1.23-alpine AS builder
LABEL org.opencontainers.image.source="https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
