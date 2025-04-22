package app

import (
	"os"
	"time"

	"site_builder_backend/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

// RunMigrations runs database migrations using GORM
func RunMigrations(l logger.Logger, models ...interface{}) {
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		l.Fatal("migrate: environment variable not declared: PG_URL")
	}

	// Create a new GORM logger adapter
	gormLogger := &logger.GormLoggerAdapter{
		Logger: l,
	}

	var (
		attempts = _defaultAttempts
		err      error
		db       *gorm.DB
	)

	for attempts > 0 {
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
			Logger: gormLogger,
		})

		if err == nil {
			break
		}

		l.Info("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		l.Fatal("Migrate: postgres connect error: %s", err)
	}

	// Run migrations for each model
	err = db.AutoMigrate(models...)
	if err != nil {
		l.Fatal("Migrate: failed to run migrations: %s", err)
	}

	l.Info("Migrate: migrations completed successfully")
}
