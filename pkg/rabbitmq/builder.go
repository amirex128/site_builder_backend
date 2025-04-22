package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ConsumerHandlerFunc defines a function that handles a RabbitMQ delivery
type ConsumerHandlerFunc func(ctx context.Context, msg amqp.Delivery) error

// Builder provides a fluent interface for configuring RabbitMQ exchanges, queues and consumers
type Builder struct {
	client        *Client
	exchangeName  string
	exchangeType  string
	queueName     string
	routingKey    string
	durable       bool
	autoDelete    bool
	exclusive     bool
	noWait        bool
	prefetchCount int
	args          amqp.Table
}

// Exchange sets the exchange name for the builder
func (c *Client) Exchange(name string) *Builder {
	return &Builder{
		client:        c,
		exchangeName:  name,
		exchangeType:  "direct", // default type
		prefetchCount: 1,        // default prefetch
	}
}

// Queue sets the queue name for the builder
func (b *Builder) Queue(name string) *Builder {
	b.queueName = name
	if b.routingKey == "" {
		b.routingKey = name // default routing key to queue name
	}
	return b
}

// Type sets the exchange type for the builder
func (b *Builder) Type(exchangeType string) *Builder {
	b.exchangeType = exchangeType
	return b
}

// RoutingKey sets the routing key for binding the queue to the exchange
func (b *Builder) RoutingKey(key string) *Builder {
	b.routingKey = key
	return b
}

// Config sets configuration options for the queue
func (b *Builder) Config(durable, autoDelete, exclusive, noWait bool) *Builder {
	b.durable = durable
	b.autoDelete = autoDelete
	b.exclusive = exclusive
	b.noWait = noWait
	return b
}

// PrefetchCount sets the prefetch count for the channel
func (b *Builder) PrefetchCount(count int) *Builder {
	b.prefetchCount = count
	return b
}

// Args sets additional arguments for queue/exchange declaration
func (b *Builder) Args(args amqp.Table) *Builder {
	b.args = args
	return b
}

// build creates all required RabbitMQ objects
func (b *Builder) build() error {
	// Declare exchange
	err := b.client.channel.ExchangeDeclare(
		b.exchangeName,
		b.exchangeType,
		b.durable,
		b.autoDelete,
		false, // internal
		b.noWait,
		b.args,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange %s: %w", b.exchangeName, err)
	}

	// Declare queue
	_, err = b.client.channel.QueueDeclare(
		b.queueName,
		b.durable,
		b.autoDelete,
		b.exclusive,
		b.noWait,
		b.args,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", b.queueName, err)
	}

	// Bind queue to exchange
	err = b.client.channel.QueueBind(
		b.queueName,
		b.routingKey,
		b.exchangeName,
		b.noWait,
		b.args,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue %s to exchange %s: %w", b.queueName, b.exchangeName, err)
	}

	// Set QoS for channel
	err = b.client.channel.Qos(b.prefetchCount, 0, false)
	if err != nil {
		return fmt.Errorf("failed to set QoS: %w", err)
	}

	return nil
}

// Consume sets up a consumer with the specified handler function
func (b *Builder) Consume(handler ConsumerHandlerFunc) error {
	// Build exchange, queue and binding
	if err := b.build(); err != nil {
		return err
	}

	// Set up consumer
	msgs, err := b.client.channel.Consume(
		b.queueName,
		"", // consumer tag (auto-generated)
		false,
		b.exclusive,
		false, // no-local
		b.noWait,
		b.args,
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	// Start consuming in a separate goroutine
	go func() {
		for msg := range msgs {
			ctx := context.Background()
			err := handler(ctx, msg)
			if err != nil {
				b.client.logger.Error("Error processing message: %v", err)
				msg.Nack(false, true) // Negative acknowledgment with requeue
			} else {
				msg.Ack(false) // Acknowledge message
			}
		}
	}()

	b.client.logger.Info("Consumer registered for queue: %s", b.queueName)
	return nil
}

// Publish publishes a message to the configured exchange
func (b *Builder) Publish(ctx context.Context, message []byte) error {
	// Build exchange, queue and binding if they don't exist
	if err := b.build(); err != nil {
		return err
	}

	// Publish message
	err := b.client.channel.PublishWithContext(
		ctx,
		b.exchangeName,
		b.routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
