// Package postgres implements postgres connection.
package postgres

import (
	"fmt"
	"site_builder_backend/pkg/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Postgres -.
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	logger       logger.Logger

	DB *gorm.DB
}

// New -.
func New(dsn string, log logger.Logger, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		logger:       log,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	var err error

	// Create GORM logger adapter
	gormLogger := &logger.GormLoggerAdapter{
		Logger: log,
	}

	for pg.connAttempts > 0 {
		pg.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})

		if err == nil {
			sqlDB, err := pg.DB.DB()
			if err != nil {
				return nil, fmt.Errorf("postgres - New - failed to get DB instance: %w", err)
			}

			// Set connection pool settings
			sqlDB.SetMaxIdleConns(pg.maxPoolSize)
			sqlDB.SetMaxOpenConns(pg.maxPoolSize)
			sqlDB.SetConnMaxLifetime(time.Hour)

			break
		}

		pg.logger.Info("Postgres is trying to connect, attempts left: %d", pg.connAttempts)
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - New - connAttempts == 0: %w", err)
	}

	return pg, nil
}

// Close -.
func (p *Postgres) Close() {
	if p.DB != nil {
		sqlDB, err := p.DB.DB()
		if err != nil {
			p.logger.Error("Failed to get SQL DB instance: %v", err)
			return
		}
		sqlDB.Close()
	}
}
