// Package rabbitmq implements rabbitmq connection.
package rabbitmq

import (
	"context"
	"fmt"
	"site_builder_backend/configs"
	"site_builder_backend/pkg/logger"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	_defaultConnectTimeout    = 5 * time.Second
	_defaultReconnectTimeout  = 5 * time.Second
	_defaultReconnectAttempts = 10
	_defaultPublishTimeout    = 2 * time.Second
	_defaultPrefetchCount     = 1
	_defaultPrefetchSize      = 0
)

// Client -.
type Client struct {
	connection        *amqp.Connection
	channel           *amqp.Channel
	logger            logger.Logger
	rmqConfig         configs.RMQ
	connectTimeout    time.Duration
	reconnectTimeout  time.Duration
	reconnectAttempts int
	queueDurable      bool
	exchangeDurable   bool
	prefetchCount     int
	prefetchSize      int
	notifyClose       chan *amqp.Error
}

// New creates a new RabbitMQ instance
func New(cfg configs.RMQ, log logger.Logger, opts ...Option) (*Client, error) {
	rmq := &Client{
		logger:            log,
		rmqConfig:         cfg,
		connectTimeout:    _defaultConnectTimeout,
		reconnectTimeout:  _defaultReconnectTimeout,
		reconnectAttempts: _defaultReconnectAttempts,
		queueDurable:      false,
		exchangeDurable:   false,
		prefetchCount:     _defaultPrefetchCount,
		prefetchSize:      _defaultPrefetchSize,
	}

	// Apply options
	for _, opt := range opts {
		opt(rmq)
	}

	// Connect to RabbitMQ
	err := rmq.connect()
	if err != nil {
		return nil, err
	}

	rmq.logger.Info("RabbitMQ connected successfully")
	return rmq, nil
}

// connect establishes a connection to RabbitMQ
func (r *Client) connect() error {
	var err error

	// Try to connect
	for i := 0; i < r.reconnectAttempts; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), r.connectTimeout)
		defer cancel()

		// Connect
		r.connection, err = amqp.Dial(r.rmqConfig.URL)
		if err == nil {
			break
		}

		r.logger.Warn("Failed to connect to RabbitMQ, attempt %d/%d: %v", i+1, r.reconnectAttempts, err)
		select {
		case <-ctx.Done():
			return fmt.Errorf("rabbitmq - connect timeout: %w", err)
		case <-time.After(r.reconnectTimeout):
			continue
		}
	}

	if err != nil {
		return fmt.Errorf("rabbitmq - connect failed after %d attempts: %w", r.reconnectAttempts, err)
	}

	// Create channel
	r.channel, err = r.connection.Channel()
	if err != nil {
		return fmt.Errorf("rabbitmq - failed to open channel: %w", err)
	}

	// Set prefetch
	err = r.channel.Qos(r.prefetchCount, r.prefetchSize, false)
	if err != nil {
		return fmt.Errorf("rabbitmq - failed to set QoS: %w", err)
	}

	// Set up notification on close
	r.notifyClose = r.connection.NotifyClose(make(chan *amqp.Error))

	// Start reconnection goroutine
	go r.handleReconnect()

	return nil
}

// handleReconnect automatically reconnects when connection is closed
func (r *Client) handleReconnect() {
	for {
		reason, ok := <-r.notifyClose
		if !ok {
			// Channel closed normally
			return
		}
		r.logger.Warn("RabbitMQ connection closed: %v", reason)

		// Try to reconnect
		for i := 0; i < r.reconnectAttempts; i++ {
			r.logger.Info("Attempting to reconnect to RabbitMQ, attempt %d/%d", i+1, r.reconnectAttempts)

			time.Sleep(r.reconnectTimeout)

			err := r.connect()
			if err == nil {
				r.logger.Info("Successfully reconnected to RabbitMQ")
				return
			}
			r.logger.Error("Failed to reconnect to RabbitMQ: %v", err)
		}
		r.logger.Error("Failed to reconnect to RabbitMQ after %d attempts", r.reconnectAttempts)
		return
	}
}

// DeclareExchange declares an exchange
func (r *Client) DeclareExchange(name, exchangeType string) error {
	return r.channel.ExchangeDeclare(
		name,              // name
		exchangeType,      // type
		r.exchangeDurable, // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
}

// DeclareQueue declares a queue
func (r *Client) DeclareQueue(name string) (amqp.Queue, error) {
	return r.channel.QueueDeclare(
		name,           // name
		r.queueDurable, // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
}

// BindQueue binds a queue to an exchange
func (r *Client) BindQueue(queueName, key, exchangeName string) error {
	return r.channel.QueueBind(
		queueName,    // queue name
		key,          // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
}

// Publish publishes a message to an exchange
func (r *Client) Publish(ctx context.Context, exchange, routingKey string, message []byte) error {
	return r.channel.PublishWithContext(
		ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}

// Consume consumes messages from a queue
func (r *Client) Consume(queueName, consumerName string) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queueName,    // queue
		consumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
}

// Close closes the RabbitMQ connection
func (r *Client) Close() error {
	if r.channel != nil {
		if err := r.channel.Close(); err != nil {
			r.logger.Error("Failed to close RabbitMQ channel: %v", err)
		}
	}

	if r.connection != nil {
		if err := r.connection.Close(); err != nil {
			return fmt.Errorf("rabbitmq - failed to close connection: %w", err)
		}
	}

	return nil
}

// ServerExchange returns the server exchange name
func (r *Client) ServerExchange() string {
	return r.rmqConfig.ServerExchange
}

// ClientExchange returns the client exchange name
func (r *Client) ClientExchange() string {
	return r.rmqConfig.ClientExchange
}

// Channel returns the AMQP channel
func (r *Client) Channel() *amqp.Channel {
	return r.channel
}

// Connection returns the AMQP connection
func (r *Client) Connection() *amqp.Connection {
	return r.connection
}
