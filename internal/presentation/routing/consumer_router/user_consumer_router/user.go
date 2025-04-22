package user_consumer_router

import (
	"site_builder_backend/internal/adapters/consumer/user_consumer"
	"site_builder_backend/internal/application/use_cases/user_use_case"
	"site_builder_backend/internal/infrastructures/impl/db/mysql/user_repo"
	"site_builder_backend/pkg/postgres"
	"site_builder_backend/pkg/rabbitmq"
)

func UserRegister(c *rabbitmq.Client, db *postgres.Postgres) {
	// Initialize repository, use case, and consumer
	repo := user_repo.NewUserRepository(db.DB)
	useCase := user_use_case.NewUserUseCase(repo)
	consumer := user_consumer.NewUserConsumer(useCase)

	// Register SMS consumer
	err := c.Exchange("user_exchange").
		Queue("sms_queue").
		Type("direct").
		RoutingKey("user.sms").
		Config(true, false, false, false).
		Consume(consumer.SendSmsConsume)

	if err != nil {
		panic("Failed to register SMS consumer: " + err.Error())
	}
}
