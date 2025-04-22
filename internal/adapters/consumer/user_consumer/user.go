package user_consumer

import (
	"context"
	"encoding/json"
	"site_builder_backend/internal/application/use_cases/user_use_case"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserConsumer struct {
	useCase *user_use_case.UserUseCase
}

func NewUserConsumer(useCase *user_use_case.UserUseCase) *UserConsumer {
	return &UserConsumer{
		useCase: useCase,
	}
}

// SendSmsConsume handles SMS notification messages from RabbitMQ
func (c *UserConsumer) SendSmsConsume(ctx context.Context, msg amqp.Delivery) error {
	// Parse message body
	var smsRequest struct {
		Phone   string `json:"phone"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal(msg.Body, &smsRequest); err != nil {
		return err
	}

	// Call use case to send SMS
	return c.useCase.SendSms(ctx, smsRequest.Phone, smsRequest.Message)
}
