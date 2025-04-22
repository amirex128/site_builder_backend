// Package redis implements redis connection.
package redis

import (
	"context"
	"fmt"
	"site_builder_backend/configs"
	"site_builder_backend/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_defaultConnectTimeout = 5 * time.Second
	_defaultReadTimeout    = 2 * time.Second
	_defaultWriteTimeout   = 2 * time.Second
)

// Redis -.
type Redis struct {
	clients        map[int]*redis.Client
	logger         logger.Logger
	redisConf      configs.Redis
	readTimeout    time.Duration
	writeTimeout   time.Duration
	connectTimeout time.Duration
	poolSize       int
	minIdleConns   int
}

// New creates a new Redis instance with multiple databases
func New(cfg configs.Redis, log logger.Logger, opts ...Option) (*Redis, error) {
	r := &Redis{
		clients:        make(map[int]*redis.Client),
		logger:         log,
		redisConf:      cfg,
		readTimeout:    _defaultReadTimeout,
		writeTimeout:   _defaultWriteTimeout,
		connectTimeout: _defaultConnectTimeout,
		poolSize:       cfg.PoolSize,
		minIdleConns:   cfg.MinIdleConns,
	}

	// Apply options
	for _, opt := range opts {
		opt(r)
	}

	// Create a list of databases to initialize
	dbList := []int{
		cfg.DefaultDB,
		cfg.CacheDB,
		cfg.SessionDB,
		cfg.RateLimiterDB,
	}

	// Initialize each database client
	for _, db := range dbList {
		client, err := r.createClient(db)
		if err != nil {
			return nil, err
		}
		r.clients[db] = client
	}

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), r.connectTimeout)
	defer cancel()

	if err := r.clients[cfg.DefaultDB].Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis - ping failed: %w", err)
	}

	r.logger.Info("Redis connected successfully")
	return r, nil
}

// GetClient returns a redis client for the specified database
func (r *Redis) GetClient(db int) (*redis.Client, error) {
	client, exists := r.clients[db]
	if !exists {
		var err error
		client, err = r.createClient(db)
		if err != nil {
			return nil, err
		}
		r.clients[db] = client
	}
	return client, nil
}

// DefaultClient returns the default redis client
func (r *Redis) DefaultClient() *redis.Client {
	return r.clients[r.redisConf.DefaultDB]
}

// CacheClient returns the cache redis client
func (r *Redis) CacheClient() *redis.Client {
	return r.clients[r.redisConf.CacheDB]
}

// SessionClient returns the session redis client
func (r *Redis) SessionClient() *redis.Client {
	return r.clients[r.redisConf.SessionDB]
}

// RateLimiterClient returns the rate limiter redis client
func (r *Redis) RateLimiterClient() *redis.Client {
	return r.clients[r.redisConf.RateLimiterDB]
}

// createClient creates a new redis client for the specified database
func (r *Redis) createClient(db int) (*redis.Client, error) {
	opts, err := redis.ParseURL(r.redisConf.URL)
	if err != nil {
		return nil, fmt.Errorf("redis - parse URL: %w", err)
	}

	// Set additional options
	opts.DB = db
	if r.redisConf.Password != "" {
		opts.Password = r.redisConf.Password
	}
	opts.PoolSize = r.poolSize
	opts.MinIdleConns = r.minIdleConns
	opts.ReadTimeout = r.readTimeout
	opts.WriteTimeout = r.writeTimeout

	client := redis.NewClient(opts)
	return client, nil
}

// Close closes all redis clients
func (r *Redis) Close() error {
	for db, client := range r.clients {
		if err := client.Close(); err != nil {
			r.logger.Error("redis - failed to close client for DB %d: %v", db, err)
		}
	}
	return nil
}
