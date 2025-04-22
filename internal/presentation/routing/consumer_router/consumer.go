package consumer_router

import (
	"site_builder_backend/internal/presentation/routing/consumer_router/user_consumer_router"
	"site_builder_backend/pkg/elasticsearch"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/postgres"
	"site_builder_backend/pkg/rabbitmq"
	"site_builder_backend/pkg/redis"
)

// Register registers all consumer routes
func Register(l *logger.ZapLogger, client *rabbitmq.Client, pgClient *postgres.Postgres, redisClient *redis.Redis, esClient *elasticsearch.Elasticsearch) {
	// User Consumer routes
	user_consumer_router.UserRegister(client, pgClient)
}
