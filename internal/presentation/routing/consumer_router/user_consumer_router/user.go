package user_consumer_router

import (
	"site_builder_backend/internal/adapters/consumer/user_consumer"
	"site_builder_backend/internal/application/use_cases/user_use_case"
	"site_builder_backend/internal/presentation/routing"
	"site_builder_backend/pkg/rabbitmq"
)

func UserRegister(client *rabbitmq.Client, services *routing.Services) {
	// Initialize repository, use case, and consumer
	useCase := user_use_case.NewUserUseCase(services.UserReadRepo, nil, nil)
	consumer := user_consumer.NewUserConsumer(useCase)

	// Register SMS consumer
	err := client.Exchange("user_exchange").
		Queue("sms_queue").
		Type("direct").
		RoutingKey("user.sms").
		Config(true, false, false, false).
		Consume(consumer.SendSmsConsume)

	if err != nil {
		panic("Failed to register SMS consumer: " + err.Error())
	}
}
