package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher provides a builder for publishing messages to RabbitMQ
type Publisher struct {
	client            *Client
	exchangeName      string
	exchangeType      string
	routingKey        string
	durable           bool
	autoDelete        bool
	internal          bool
	noWait            bool
	args              amqp.Table
	contentType       string
	deliveryMode      uint8
	expiration        string
	priority          uint8
	publisherConfirms bool
	confirmsChannel   chan amqp.Confirmation
}

// NewPublisher creates a new publisher builder
func (c *Client) Publisher(exchangeName string) *Publisher {
	return &Publisher{
		client:            c,
		exchangeName:      exchangeName,
		exchangeType:      "direct", // default type
		contentType:       "application/json",
		deliveryMode:      amqp.Persistent,
		publisherConfirms: false,
	}
}

// Type sets the exchange type for the publisher
func (p *Publisher) Type(exchangeType string) *Publisher {
	p.exchangeType = exchangeType
	return p
}

// RoutingKey sets the routing key for publishing messages
func (p *Publisher) RoutingKey(key string) *Publisher {
	p.routingKey = key
	return p
}

// Config sets configuration options for the exchange
func (p *Publisher) Config(durable, autoDelete, internal, noWait bool) *Publisher {
	p.durable = durable
	p.autoDelete = autoDelete
	p.internal = internal
	p.noWait = noWait
	return p
}

// ContentType sets the content type for messages
func (p *Publisher) ContentType(contentType string) *Publisher {
	p.contentType = contentType
	return p
}

// DeliveryMode sets the delivery mode (persistent or transient)
func (p *Publisher) DeliveryMode(persistent bool) *Publisher {
	if persistent {
		p.deliveryMode = amqp.Persistent
	} else {
		p.deliveryMode = amqp.Transient
	}
	return p
}

// Expiration sets the expiration time for messages
func (p *Publisher) Expiration(expiration string) *Publisher {
	p.expiration = expiration
	return p
}

// Priority sets the priority for messages
func (p *Publisher) Priority(priority uint8) *Publisher {
	p.priority = priority
	return p
}

// WithConfirms enables publisher confirms
func (p *Publisher) WithConfirms() *Publisher {
	p.publisherConfirms = true
	return p
}

// Args sets additional arguments for exchange declaration
func (p *Publisher) Args(args amqp.Table) *Publisher {
	p.args = args
	return p
}

// build creates the exchange if it doesn't exist
func (p *Publisher) build() error {
	// Declare exchange
	err := p.client.channel.ExchangeDeclare(
		p.exchangeName,
		p.exchangeType,
		p.durable,
		p.autoDelete,
		p.internal,
		p.noWait,
		p.args,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange %s: %w", p.exchangeName, err)
	}

	// Enable publisher confirms if requested
	if p.publisherConfirms {
		if err := p.client.channel.Confirm(false); err != nil {
			return fmt.Errorf("failed to put channel in confirm mode: %w", err)
		}
		p.confirmsChannel = p.client.channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	}

	return nil
}

// Publish publishes a message to the exchange
func (p *Publisher) Publish(ctx context.Context, message []byte) error {
	// Build exchange if needed
	if err := p.build(); err != nil {
		return err
	}

	// Publish message
	err := p.client.channel.PublishWithContext(
		ctx,
		p.exchangeName,
		p.routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  p.contentType,
			DeliveryMode: p.deliveryMode,
			Expiration:   p.expiration,
			Priority:     p.priority,
			Timestamp:    time.Now(),
			Body:         message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	// Wait for confirmation if enabled
	if p.publisherConfirms {
		select {
		case confirm := <-p.confirmsChannel:
			if !confirm.Ack {
				return fmt.Errorf("publisher confirms received nack")
			}
		case <-ctx.Done():
			return fmt.Errorf("publisher confirms timed out: %w", ctx.Err())
		}
	}

	return nil
}

// PublishJSON marshals a struct to JSON and publishes it
func (p *Publisher) PublishJSON(ctx context.Context, payload interface{}) error {
	// Marshal payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Set content type to JSON if not explicitly changed
	if p.contentType == "" {
		p.contentType = "application/json"
	}

	// Publish the message
	return p.Publish(ctx, jsonData)
}

// PublishBatch publishes multiple messages in a batch
func (p *Publisher) PublishBatch(ctx context.Context, messages [][]byte) error {
	// Build exchange if needed
	if err := p.build(); err != nil {
		return err
	}

	// Publish all messages
	for _, message := range messages {
		err := p.client.channel.PublishWithContext(
			ctx,
			p.exchangeName,
			p.routingKey,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType:  p.contentType,
				DeliveryMode: p.deliveryMode,
				Expiration:   p.expiration,
				Priority:     p.priority,
				Timestamp:    time.Now(),
				Body:         message,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to publish message in batch: %w", err)
		}
	}

	return nil
}
