// internal/infrastructure/server/server.go
package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/routes"
)

type Server struct {
	router *routes.Router
	config *config.Config
	logger *slog.Logger
}

func NewServer(cfg *config.Config, logger *slog.Logger, handlers *routes.Handlers) *Server {
	// Cria o router
	router := routes.NewRouter(logger, cfg.Environment)

	// Registra as rotas
	router.RegisterRoutes(handlers)

	return &Server{
		router: router,
		config: cfg,
		logger: logger,
	}
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:         ":" + s.config.ServerPort,
		Handler:      s.router.Engine(),
		ReadTimeout:  s.config.ServerReadTimeout,
		WriteTimeout: s.config.ServerWriteTimeout,
		IdleTimeout:  s.config.ServerIdleTimeout,
	}

	// Graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		s.logger.Info("shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			s.logger.Error("server forced to shutdown", "error", err)
			os.Exit(1)
		}

		s.logger.Info("server exited gracefully")
	}()

	s.logger.Info("server is running", "port", s.config.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
