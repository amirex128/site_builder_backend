package rabbitmq

import "time"

// Option -.
type Option func(*Client)

// ConnectTimeout -.
func ConnectTimeout(timeout time.Duration) Option {
	return func(r *Client) {
		r.connectTimeout = timeout
	}
}

// ReconnectAttempts -.
func ReconnectAttempts(attempts int) Option {
	return func(r *Client) {
		r.reconnectAttempts = attempts
	}
}

// ReconnectTimeout -.
func ReconnectTimeout(timeout time.Duration) Option {
	return func(r *Client) {
		r.reconnectTimeout = timeout
	}
}

// QueueDurable -.
func QueueDurable(durable bool) Option {
	return func(r *Client) {
		r.queueDurable = durable
	}
}

// ExchangeDurable -.
func ExchangeDurable(durable bool) Option {
	return func(r *Client) {
		r.exchangeDurable = durable
	}
}

// PrefetchCount -.
func PrefetchCount(count int) Option {
	return func(r *Client) {
		r.prefetchCount = count
	}
}

// PrefetchSize -.
func PrefetchSize(size int) Option {
	return func(r *Client) {
		r.prefetchSize = size
	}
}
