package main

import (
	"site_builder_backend/configs"
	"site_builder_backend/internal/presentation/app"
	"site_builder_backend/pkg/logger"
)

func main() {
	// Initialize basic logger for startup
	startupLogger := logger.NewLoggerFromConfig("info", "console", "stdout")

	// Configuration
	cfg, err := configs.NewConfig()
	if err != nil {
		startupLogger.Fatal("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
