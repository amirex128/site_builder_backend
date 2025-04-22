package configs

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	// Config -.
	Config struct {
		App           App
		HTTP          HTTP
		Log           Log
		PG            PG
		RMQ           RMQ
		Redis         Redis
		Elasticsearch Elasticsearch
		Metrics       Metrics
		Swagger       Swagger
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	// Log -.
	Log struct {
		Level  string `env:"LOG_LEVEL,required"`
		Format string `env:"LOG_FORMAT" envDefault:"json"`
		Output string `env:"LOG_OUTPUT" envDefault:"stdout"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT,required"`
		URL            string `env:"RMQ_URL,required"`
	}

	// Redis - Multiple database configuration
	Redis struct {
		URL           string `env:"REDIS_URL" envDefault:"redis://localhost:6379"`
		PoolSize      int    `env:"REDIS_POOL_SIZE" envDefault:"10"`
		MinIdleConns  int    `env:"REDIS_MIN_IDLE_CONNS" envDefault:"5"`
		DefaultDB     int    `env:"REDIS_DEFAULT_DB" envDefault:"0"`
		CacheDB       int    `env:"REDIS_CACHE_DB" envDefault:"1"`
		SessionDB     int    `env:"REDIS_SESSION_DB" envDefault:"2"`
		RateLimiterDB int    `env:"REDIS_RATE_LIMITER_DB" envDefault:"3"`
		Password      string `env:"REDIS_PASSWORD" envDefault:""`
	}

	// Elasticsearch - Configuration
	Elasticsearch struct {
		Addresses      []string `env:"ELASTICSEARCH_ADDRESSES" envDefault:"http://localhost:9200"`
		Username       string   `env:"ELASTICSEARCH_USERNAME" envDefault:""`
		Password       string   `env:"ELASTICSEARCH_PASSWORD" envDefault:""`
		EnabledIndexes []string `env:"ELASTICSEARCH_ENABLED_INDEXES" envDefault:"products,blogs"`
		SniffEnabled   bool     `env:"ELASTICSEARCH_SNIFF" envDefault:"false"`
	}

	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"true"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("config - NewConfig - env.Parse: %w", err)
	}

	return cfg, nil
}
