package examples

import (
	"context"
	"encoding/json"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/rabbitmq"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// SmsHandler handles SMS messages from RabbitMQ
type SmsHandler struct {
	logger logger.Logger
}

// NewSmsHandler creates a new SMS handler
func NewSmsHandler(log logger.Logger) *SmsHandler {
	return &SmsHandler{
		logger: log,
	}
}

// HandleSms processes SMS messages
func (h *SmsHandler) HandleSms(ctx context.Context, msg amqp.Delivery) error {
	var smsMessage SmsMessage
	if err := json.Unmarshal(msg.Body, &smsMessage); err != nil {
		h.logger.Error("Failed to unmarshal SMS message: %v", err)
		return err
	}

	h.logger.Info("Processing SMS message: To=%s, Message=%s, Type=%s",
		smsMessage.Phone, smsMessage.Message, smsMessage.Type)

	// Simulate processing time
	time.Sleep(100 * time.Millisecond)

	// In a real application, you would call an SMS service here
	// For example:
	// return sendSmsToProvider(smsMessage.Phone, smsMessage.Message)

	return nil // success
}

// CompleteExample demonstrates a complete example with both consumer and publisher
func CompleteExample(client *rabbitmq.Client, log logger.Logger) {
	// Create the handler
	handler := NewSmsHandler(log)

	// Set up the exchange, queue, and consumer for SMS messages
	err := client.Exchange("user_exchange").
		Queue("sms_queue").
		Type("direct").
		RoutingKey("user.sms").
		Config(true, false, false, false). // durable, autoDelete, exclusive, noWait
		Consume(handler.HandleSms)

	if err != nil {
		log.Error("Failed to set up SMS consumer: %v", err)
		return
	}

	log.Info("SMS consumer registered successfully")

	// Wait a moment for the consumer to be ready
	time.Sleep(500 * time.Millisecond)

	// Now publish a message to be consumed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	smsMessage := SmsMessage{
		Phone:   "+123456789",
		Message: "Hello from the complete example!",
		Type:    "example",
	}

	// Publish using the publisher builder
	err = client.Publisher("user_exchange").
		RoutingKey("user.sms").
		PublishJSON(ctx, smsMessage)

	if err != nil {
		log.Error("Failed to publish example message: %v", err)
		return
	}

	log.Info("Example message published successfully")

	// Keep the application running to demonstrate the message being consumed
	log.Info("Waiting for message to be processed...")
	time.Sleep(1 * time.Second)
}

// DemoPublishAndConsume shows how to build a topology with multiple consumers and publishers
func DemoPublishAndConsume(client *rabbitmq.Client, log logger.Logger) {
	// Create the handler
	handler := NewSmsHandler(log)

	// Set up multiple consumers for different message types
	setupConsumers(client, handler, log)

	// Publish different types of messages
	publishMessages(client, log)

	// Wait for messages to be processed
	log.Info("Waiting for messages to be processed...")
	time.Sleep(2 * time.Second)
}

func setupConsumers(client *rabbitmq.Client, handler *SmsHandler, log logger.Logger) {
	// Set up a consumer for verification messages
	err := client.Exchange("user_exchange").
		Queue("verification_queue").
		Type("topic").
		RoutingKey("user.sms.verification").
		PrefetchCount(5).
		Consume(handler.HandleSms)

	if err != nil {
		log.Error("Failed to set up verification consumer: %v", err)
	} else {
		log.Info("Verification consumer registered successfully")
	}

	// Set up a consumer for notification messages
	err = client.Exchange("user_exchange").
		Queue("notification_queue").
		Type("topic").
		RoutingKey("user.sms.notification").
		PrefetchCount(10).
		Consume(handler.HandleSms)

	if err != nil {
		log.Error("Failed to set up notification consumer: %v", err)
	} else {
		log.Info("Notification consumer registered successfully")
	}

	// Set up a consumer for all SMS messages (using wildcard)
	err = client.Exchange("user_exchange").
		Queue("all_sms_queue").
		Type("topic").
		RoutingKey("user.sms.*").
		Consume(handler.HandleSms)

	if err != nil {
		log.Error("Failed to set up wildcard consumer: %v", err)
	} else {
		log.Info("Wildcard consumer registered successfully")
	}
}

func publishMessages(client *rabbitmq.Client, log logger.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish a verification message
	verificationMsg := SmsMessage{
		Phone:   "+111222333",
		Message: "Your verification code is 54321",
		Type:    "verification",
	}

	err := client.Publisher("user_exchange").
		RoutingKey("user.sms.verification").
		Type("topic").
		Config(true, false, false, false).
		PublishJSON(ctx, verificationMsg)

	if err != nil {
		log.Error("Failed to publish verification message: %v", err)
	} else {
		log.Info("Verification message published successfully")
	}

	// Publish a notification message
	notificationMsg := SmsMessage{
		Phone:   "+444555666",
		Message: "You have a new message from John",
		Type:    "notification",
	}

	err = client.Publisher("user_exchange").
		RoutingKey("user.sms.notification").
		Type("topic").
		Config(true, false, false, false).
		PublishJSON(ctx, notificationMsg)

	if err != nil {
		log.Error("Failed to publish notification message: %v", err)
	} else {
		log.Info("Notification message published successfully")
	}
}
