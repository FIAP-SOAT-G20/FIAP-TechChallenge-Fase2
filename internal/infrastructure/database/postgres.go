package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/infrastructure/config"
)

type Database struct {
	*gorm.DB
}

type GormLogger struct {
	*slog.Logger
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.InfoContext(ctx, msg, "data", data)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.WarnContext(ctx, msg, "data", data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.ErrorContext(ctx, msg, "data", data)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		l.Logger.ErrorContext(ctx, "database query failed",
			"error", err,
			"elapsed", elapsed,
			"sql", sql,
			"rows", rows,
		)
		return
	}

	l.Logger.DebugContext(ctx, "database query",
		"elapsed", elapsed,
		"sql", sql,
		"rows", rows,
	)
}

func NewPostgresConnection(cfg *config.Config, logger *slog.Logger) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	// Configure GORM with slog logger
	gormConfig := &gorm.Config{
		Logger: &GormLogger{Logger: logger},
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused
	sqlDB.SetConnMaxLifetime(cfg.DBMaxLifetime)

	return &Database{db}, nil
}

// Migrate runs database migrations
func (db *Database) Migrate() error {
	return db.AutoMigrate(
		&entity.Product{},
	)
}
