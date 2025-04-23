package routing

import (
	"github.com/gin-gonic/gin"
	"site_builder_backend/pkg/elasticsearch"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/postgres"
	"site_builder_backend/pkg/redis"
)

type Services struct {
	Logger         *logger.ZapLogger
	ElasticClient  *elasticsearch.Elasticsearch
	RedisClient    *redis.Redis
	PostgresClient *postgres.Postgres
	Router         *gin.Engine
}

func NewServices(router *gin.Engine, pgClient *postgres.Postgres, redisClient *redis.Redis, esClient *elasticsearch.Elasticsearch, l *logger.ZapLogger) *Services {
	return &Services{
		Router:         router,
		Logger:         l,
		ElasticClient:  esClient,
		RedisClient:    redisClient,
		PostgresClient: pgClient,
	}
}
