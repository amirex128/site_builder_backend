package consumer_router

import (
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ConsumerHandler handles message consumption
type ConsumerHandler struct {
	logger         logger.Logger
	rabbitmqClient *rabbitmq.Client
	consumers      map[string]*rabbitmq.Consumer
}

// NewConsumerHandler creates a new consumer handler
func NewConsumerHandler(logger logger.Logger, rabbitmqClient *rabbitmq.Client) *ConsumerHandler {
	return &ConsumerHandler{
		logger:         logger,
		rabbitmqClient: rabbitmqClient,
		consumers:      make(map[string]*rabbitmq.Consumer),
	}
}

// RegisterConsumers registers all consumers
func (h *ConsumerHandler) RegisterConsumers() error {
	h.logger.Info("Registering consumers...")

	// Register exchange for consumers
	err := h.rabbitmqClient.DeclareExchange(h.rabbitmqClient.ServerExchange(), "direct")
	if err != nil {
		h.logger.Error("Failed to declare server exchange: %v", err)
		return err
	}

	// Register your consumers here
	err = h.registerExampleConsumer()
	if err != nil {
		return err
	}
	// Add more consumers as needed:
	// h.registerOrderConsumer()
	// h.registerNotificationConsumer()

	return nil
}

// Example consumer implementation
func (h *ConsumerHandler) registerExampleConsumer() error {
	queueName := "example_queue"

	// Create a new consumer
	consumer := rabbitmq.NewConsumer(h.rabbitmqClient, h.logger, queueName)

	// Register handlers for different routing keys
	err := consumer.RegisterHandler("example.key", h.handleExampleMessage)
	if err != nil {
		h.logger.Error("Failed to register handler: %v", err)
		return err
	}

	// Start the consumer
	err = consumer.Start()
	if err != nil {
		h.logger.Error("Failed to start consumer: %v", err)
		return err
	}

	// Store the consumer for potential later use
	h.consumers[queueName] = consumer

	h.logger.Info("Example consumer registered")
	return nil
}

// Handler for example messages
func (h *ConsumerHandler) handleExampleMessage(msg amqp.Delivery) error {
	h.logger.Info("Received example message: %s", string(msg.Body))

	// Process message
	// ...

	return nil // Return nil for success, error for failure
}

// Register sets up the message consumer
func Register(logger logger.Logger, rabbitmqClient *rabbitmq.Client) {
	handler := NewConsumerHandler(logger, rabbitmqClient)
	err := handler.RegisterConsumers()
	if err != nil {
		logger.Error("Failed to register consumers: %v", err)
	}
}
