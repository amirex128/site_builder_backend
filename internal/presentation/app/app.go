package app

import (
	"os"
	"os/signal"
	"site_builder_backend/configs"
	"site_builder_backend/internal/presentation/middlewares"
	"site_builder_backend/internal/presentation/routing/consumer_router"
	"site_builder_backend/internal/presentation/routing/http_router"
	"site_builder_backend/pkg/elasticsearch"
	"site_builder_backend/pkg/httpserver"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/postgres"
	"site_builder_backend/pkg/rabbitmq"
	"site_builder_backend/pkg/redis"
	"syscall"
	"time"
)

const (
	_defaultConnectTimeout = 5 * time.Second
)

// Run starts the application
func Run(cfg *configs.Config) {
	// Initialize logger with Zap
	l := logger.NewLogger(logger.LogConfig{
		Level:  cfg.Log.Level,
		Format: cfg.Log.Format,
		Output: cfg.Log.Output,
	})

	var pgClient *postgres.Postgres
	var redisClient *redis.Redis
	var esClient *elasticsearch.Elasticsearch
	var rmqClient *rabbitmq.Client
	var err error

	// Initialize PostgreSQL with GORM
	pgClient, err = postgres.New(cfg.PG.URL, l, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal("app - Run - postgres.New: %v", err)
	}
	defer func() {
		if pgClient != nil {
			pgClient.Close()
		}
	}()

	// Initialize Redis with multiple databases
	redisClient, err = redis.New(cfg.Redis, l,
		redis.ConnectTimeout(_defaultConnectTimeout),
		redis.PoolSize(cfg.Redis.PoolSize),
		redis.MinIdleConns(cfg.Redis.MinIdleConns))
	if err != nil {
		l.Fatal("app - Run - redis.New: %v", err)
	}
	defer func() {
		if redisClient != nil {
			redisClient.Close()
		}
	}()

	// Initialize Elasticsearch
	esClient, err = elasticsearch.New(cfg.Elasticsearch, l,
		elasticsearch.ConnectTimeout(_defaultConnectTimeout),
		elasticsearch.SniffEnabled(cfg.Elasticsearch.SniffEnabled),
		elasticsearch.Retries(3))
	if err != nil {
		l.Fatal("app - Run - elasticsearch.New: %v", err)
	}
	// Log enabled Elasticsearch indexes
	l.Info("Elasticsearch initialized with indexes: %v", esClient.EnabledIndexes())

	// Initialize RabbitMQ
	rmqClient, err = rabbitmq.New(cfg.RMQ, l,
		rabbitmq.ConnectTimeout(_defaultConnectTimeout),
		rabbitmq.ReconnectAttempts(5),
		rabbitmq.QueueDurable(true),
		rabbitmq.ExchangeDurable(true))
	if err != nil {
		l.Fatal("app - Run - rabbitmq.New: %v", err)
	}
	defer func() {
		if rmqClient != nil {
			rmqClient.Close()
		}
	}()

	// HTTP Server with Gin
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))

	// Register middleware
	httpServer.App.Use(middlewares.LoggerMiddleware(l))
	httpServer.App.Use(middlewares.RecoveryMiddleware(l))

	// Initialize HTTP routers
	http_router.Register(l, httpServer.App, pgClient, redisClient, esClient)

	// Initialize message consumers
	consumer_router.Register(l, rmqClient, pgClient, redisClient, esClient)

	// Start server
	httpServer.Start()
	l.Info("HTTP server started on port %s", cfg.HTTP.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal received: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error("app - Run - httpServer.Notify: %v", err)
	}

	// Shutdown
	l.Info("Shutting down server...")
	err = httpServer.Shutdown()
	if err != nil {
		l.Error("app - Run - httpServer.Shutdown: %v", err)
	}
}
