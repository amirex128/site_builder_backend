package app

import (
	"os"
	"os/signal"
	"site_builder_backend/configs"
	"site_builder_backend/internal/presentation/middlewares"
	"site_builder_backend/internal/presentation/routing"
	"site_builder_backend/internal/presentation/routing/consumer_router"
	"site_builder_backend/internal/presentation/routing/http_router"
	"site_builder_backend/pkg/httpserver"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/rabbitmq"
	"syscall"
)

// Run starts the application
func Run(cfg *configs.Config) {
	l := logger.NewLogger(logger.LogConfig{
		Level:  cfg.Log.Level,
		Format: cfg.Log.Format,
		Output: cfg.Log.Output,
	})
	rmqClient, err := rabbitmq.New(cfg.RMQ, l,
		rabbitmq.ConnectTimeout(3000),
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
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))

	httpServer.App.Use(middlewares.LoggerMiddleware(l))
	httpServer.App.Use(middlewares.RecoveryMiddleware(l))

	services := routing.NewServiceRegistration(cfg, l)
	controllerServices := routing.NewControllerServices(services)

	http_router.Register(httpServer.App, controllerServices, services)

	consumer_router.Register(rmqClient, services)

	httpServer.Start()
	l.Info("HTTP server started on port %s", cfg.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal received: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error("app - Run - httpServer.Notify: %v", err)
	}

	l.Info("Shutting down server...")
	err = httpServer.Shutdown()
	if err != nil {
		l.Error("app - Run - httpServer.Shutdown: %v", err)
	}
}
