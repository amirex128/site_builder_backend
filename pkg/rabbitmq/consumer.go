package rabbitmq

import (
	"fmt"
	"site_builder_backend/pkg/logger"

	amqp "github.com/rabbitmq/amqp091-go"
)

// MessageHandler defines a function that processes a RabbitMQ message
type MessageHandler func(msg amqp.Delivery) error

// Consumer defines a RabbitMQ consumer structure
type Consumer struct {
	client    *Client
	logger    logger.Logger
	queue     string
	handlers  map[string]MessageHandler
	consumers map[string]<-chan amqp.Delivery
}

// NewConsumer creates a new RabbitMQ consumer
func NewConsumer(client *Client, log logger.Logger, queue string) *Consumer {
	return &Consumer{
		client:    client,
		logger:    log,
		queue:     queue,
		handlers:  make(map[string]MessageHandler),
		consumers: make(map[string]<-chan amqp.Delivery),
	}
}

// RegisterHandler registers a handler for a specific routing key
func (c *Consumer) RegisterHandler(routingKey string, handler MessageHandler) error {
	if _, exists := c.handlers[routingKey]; exists {
		return fmt.Errorf("handler for routing key %s already registered", routingKey)
	}
	c.handlers[routingKey] = handler
	return nil
}

// Start begins consuming messages from the queue
func (c *Consumer) Start() error {
	// Declare the queue if it doesn't exist
	_, err := c.client.DeclareQueue(c.queue)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Set up bindings for each routing key
	for routingKey := range c.handlers {
		err = c.client.BindQueue(c.queue, routingKey, c.client.ServerExchange())
		if err != nil {
			return fmt.Errorf("failed to bind queue to routing key %s: %w", routingKey, err)
		}
	}

	// Start consuming
	msgs, err := c.client.Consume(c.queue, "")
	if err != nil {
		return fmt.Errorf("failed to consume from queue: %w", err)
	}

	// Process messages in a goroutine
	go c.processMessages(msgs)

	c.logger.Info("Consumer started for queue: %s", c.queue)
	return nil
}

// processMessages handles incoming messages and routes them to appropriate handlers
func (c *Consumer) processMessages(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		routingKey := msg.RoutingKey
		handler, exists := c.handlers[routingKey]

		if !exists {
			c.logger.Warn("No handler registered for routing key: %s", routingKey)
			msg.Reject(false)
			continue
		}

		// Process the message
		err := handler(msg)
		if err != nil {
			c.logger.Error("Failed to process message with routing key %s: %v", routingKey, err)
			msg.Reject(true) // Reject with requeue
		} else {
			msg.Ack(false)
		}
	}

	c.logger.Warn("Message channel closed, consumer for queue %s stopped", c.queue)
}
