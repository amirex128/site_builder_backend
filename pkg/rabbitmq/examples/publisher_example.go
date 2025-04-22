package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/rabbitmq"
	"time"
)

// SmsMessage represents an SMS message to be sent
type SmsMessage struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

// PublisherExample demonstrates how to use the Publisher builder
func PublisherExample(client *rabbitmq.Client, log logger.Logger) {
	// Basic message publishing
	basicPublish(client, log)

	// JSON message publishing
	jsonPublish(client, log)

	// Advanced publishing with confirms
	advancedPublish(client, log)

	// Batch publishing
	batchPublish(client, log)
}

func basicPublish(client *rabbitmq.Client, log logger.Logger) {
	// Create a message
	message := []byte(`{"phone": "+123456789", "message": "Hello user!", "type": "welcome"}`)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish the message
	err := client.Publisher("user_exchange").
		RoutingKey("user.notifications").
		Type("direct").
		Config(true, false, false, false). // durable, autoDelete, internal, noWait
		Publish(ctx, message)

	if err != nil {
		log.Error("Failed to publish message: %v", err)
		return
	}

	log.Info("Basic message published successfully")
}

func jsonPublish(client *rabbitmq.Client, log logger.Logger) {
	// Create a message struct
	smsMessage := SmsMessage{
		Phone:   "+987654321",
		Message: "Your verification code is 12345",
		Type:    "verification",
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish the JSON message
	err := client.Publisher("user_exchange").
		RoutingKey("user.sms").
		PublishJSON(ctx, smsMessage)

	if err != nil {
		log.Error("Failed to publish JSON message: %v", err)
		return
	}

	log.Info("JSON message published successfully")
}

func advancedPublish(client *rabbitmq.Client, log logger.Logger) {
	// Create a message
	message := []byte(`{"phone": "+123456789", "message": "Your order #12345 has been processed", "type": "order"}`)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish with advanced options
	err := client.Publisher("user_exchange").
		RoutingKey("user.important").
		Type("direct").
		Config(true, false, false, false). // durable, autoDelete, internal, noWait
		DeliveryMode(true).                // persistent
		Priority(5).                       // medium priority
		Expiration("3600000").             // expires after 1 hour
		WithConfirms().                    // enable publisher confirms
		Publish(ctx, message)

	if err != nil {
		log.Error("Failed to publish advanced message: %v", err)
		return
	}

	log.Info("Advanced message published successfully with confirms")
}

func batchPublish(client *rabbitmq.Client, log logger.Logger) {
	// Create multiple messages
	messages := make([][]byte, 0, 3)

	// Create 3 different messages
	for i := 1; i <= 3; i++ {
		smsMessage := SmsMessage{
			Phone:   fmt.Sprintf("+%d", 10000000+i),
			Message: fmt.Sprintf("Batch message #%d", i),
			Type:    "batch",
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(smsMessage)
		if err != nil {
			log.Error("Failed to marshal message: %v", err)
			continue
		}

		messages = append(messages, jsonData)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish messages in batch
	err := client.Publisher("user_exchange").
		RoutingKey("user.batch").
		Type("direct").
		Config(true, false, false, false).
		PublishBatch(ctx, messages)

	if err != nil {
		log.Error("Failed to publish batch messages: %v", err)
		return
	}

	log.Info("Batch of %d messages published successfully", len(messages))
}
