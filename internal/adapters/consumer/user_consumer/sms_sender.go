package user_consumer

import (
	"context"
	"encoding/json"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/rabbitmq"
)

// SmsSender is responsible for sending SMS messages via RabbitMQ
type SmsSender struct {
	client *rabbitmq.Client
	logger logger.Logger
}

// NewSmsSender creates a new SMS sender
func NewSmsSender(client *rabbitmq.Client, logger logger.Logger) *SmsSender {
	return &SmsSender{
		client: client,
		logger: logger,
	}
}

// SmsRequest represents a request to send an SMS
type SmsRequest struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

// SendSms sends an SMS message via RabbitMQ
func (s *SmsSender) SendSms(ctx context.Context, phone, message string) error {
	// Create the SMS request
	request := SmsRequest{
		Phone:   phone,
		Message: message,
	}

	// Publish the message using the builder pattern
	err := s.client.Publisher("user_exchange").
		RoutingKey("user.sms").
		Type("direct").
		Config(true, false, false, false). // durable, autoDelete, internal, noWait
		PublishJSON(ctx, request)

	if err != nil {
		s.logger.Error("Failed to publish SMS message: %v", err)
		return err
	}

	s.logger.Info("SMS message published successfully to %s", phone)
	return nil
}

// SendBulkSms sends multiple SMS messages at once
func (s *SmsSender) SendBulkSms(ctx context.Context, messages []SmsRequest) error {
	// Convert messages to JSON
	var jsonMessages [][]byte
	for _, msg := range messages {
		jsonData, err := json.Marshal(msg)
		if err != nil {
			s.logger.Error("Failed to marshal SMS message: %v", err)
			continue
		}
		jsonMessages = append(jsonMessages, jsonData)
	}

	// Skip if no valid messages
	if len(jsonMessages) == 0 {
		return nil
	}

	// Publish messages in batch
	err := s.client.Publisher("user_exchange").
		RoutingKey("user.sms.bulk").
		Type("direct").
		Config(true, false, false, false).
		PublishBatch(ctx, jsonMessages)

	if err != nil {
		s.logger.Error("Failed to publish bulk SMS messages: %v", err)
		return err
	}

	s.logger.Info("Bulk SMS messages published successfully (%d messages)", len(jsonMessages))
	return nil
} 