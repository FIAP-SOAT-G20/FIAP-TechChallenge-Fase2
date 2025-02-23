package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database settings
	DBHost         string
	DBPort         int
	DBUser         string
	DBPassword     string
	DBName         string
	DBMaxOpenConns int
	DBMaxIdleConns int
	DBMaxLifetime  time.Duration

	// Server settings
	ServerPort                    string
	ServerReadTimeout             time.Duration
	ServerWriteTimeout            time.Duration
	ServerIdleTimeout             time.Duration
	ServerGracefulShutdownTimeout time.Duration

	// Environment
	Environment string

	// Mercado Pago
	MercadoPagoToken           string
	MercadoPagoURL             string
	MercadoPagoTimeout         time.Duration
	MercadoPagoNotificationURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	dbMaxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
	dbMaxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "25"))
	dbMaxLifetime, _ := time.ParseDuration(getEnv("DB_CONN_MAX_LIFETIME", "5m"))

	serverReadTimeout, _ := time.ParseDuration(getEnv("SERVER_READ_TIMEOUT", "10s"))
	serverWriteTimeout, _ := time.ParseDuration(getEnv("SERVER_WRITE_TIMEOUT", "10s"))
	serverIdleTimeout, _ := time.ParseDuration(getEnv("SERVER_IDLE_TIMEOUT", "60s"))
	serverGracefulShutdownTimeout, _ := time.ParseDuration(getEnv("SERVER_GRACEFUL_SHUTDOWN_SEC_TIMEOUT", "5s"))
	mercadoPagoTimeout, _ := time.ParseDuration(getEnv("MERCADO_PAGO_TIMEOUT", "10s"))

	return &Config{
		// Database settings
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         dbPort,
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "fastfood_10soat_g18_tc2"),
		DBMaxOpenConns: dbMaxOpenConns,
		DBMaxIdleConns: dbMaxIdleConns,
		DBMaxLifetime:  dbMaxLifetime,

		// Server settings
		ServerPort:                    getEnv("SERVER_PORT", "8080"),
		ServerReadTimeout:             serverReadTimeout,
		ServerWriteTimeout:            serverWriteTimeout,
		ServerIdleTimeout:             serverIdleTimeout,
		ServerGracefulShutdownTimeout: serverGracefulShutdownTimeout,

		// Environment
		Environment: getEnv("ENVIRONMENT", "development"),

		// Mercado Pago
		MercadoPagoToken:           getEnv("MERCADO_PAGO_TOKEN", "token"),
		MercadoPagoURL:             getEnv("MERCADO_PAGO_URL", "url"),
		MercadoPagoTimeout:         mercadoPagoTimeout,
		MercadoPagoNotificationURL: getEnv("MERCADO_PAGO_NOTIFICATION_URL", "url"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
