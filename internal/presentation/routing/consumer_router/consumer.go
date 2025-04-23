package consumer_router

import (
	"site_builder_backend/internal/presentation/routing"
	"site_builder_backend/internal/presentation/routing/consumer_router/user_consumer_router"
	"site_builder_backend/pkg/rabbitmq"
)

// Register registers all consumer routes
func Register(rbClient *rabbitmq.Client, services *routing.Services) {
	user_consumer_router.UserRegister(rbClient, services)
}
