package redis

import "time"

// Option -.
type Option func(*Redis)

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.readTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.writeTimeout = timeout
	}
}

// ConnectTimeout -.
func ConnectTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.connectTimeout = timeout
	}
}

// PoolSize -.
func PoolSize(size int) Option {
	return func(r *Redis) {
		r.poolSize = size
	}
}

// MinIdleConns -.
func MinIdleConns(conns int) Option {
	return func(r *Redis) {
		r.minIdleConns = conns
	}
}
