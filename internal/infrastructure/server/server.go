package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/route"
)

type Server struct {
	router *route.Router
	config *config.Config
	logger *slog.Logger
}

func NewServer(cfg *config.Config, logger *slog.Logger, handlers *route.Handlers) *Server {
	router := route.NewRouter(logger, cfg.Environment)

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

	chanError := make(chan error)

	go gracefullyShutdown(server, chanError, s)

	if err := <-chanError; err != nil {
		s.logger.Error("server failed to start", "error", err)
	}

	return nil
}

func gracefullyShutdown(server *http.Server, chanError chan error, s *Server) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		s.logger.Info("shutting down server...")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), s.config.ServerGracefulShutdownTimeout)

		defer func() {
			stop()
			cancel()
			close(chanError)
		}()

		err := server.Shutdown(ctxTimeout)
		if err != nil {
			return
		}

		s.logger.Info("server exited gracefully")
		fmt.Println("Bye! 👋")
	}()

	go func() {
		s.logger.Info("server is running", "port", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			chanError <- fmt.Errorf("failed to start server: %w", err)
			return
		}
	}()
}
