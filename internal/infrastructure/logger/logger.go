package logger

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(environment string) *Logger {
	var handler slog.Handler

	if environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return &Logger{
		Logger: slog.New(handler),
	}
}

// WithContext adds request context information to the logger
func (l *Logger) WithContext(ctx context.Context) *slog.Logger {
	requestID, ok := ctx.Value("request_id").(string)
	if !ok {
		return l.Logger
	}

	return l.Logger.With(
		slog.String("request_id", requestID),
	)
}

// WithFields adds structured fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) *slog.Logger {
	attrs := make([]any, 0, len(fields)*2)
	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}
	return l.Logger.With(attrs...)
}
