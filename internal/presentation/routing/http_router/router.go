package http_router

import (
	"site_builder_backend/internal/presentation/routing/http_router/user_router"
	"site_builder_backend/pkg/elasticsearch"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/postgres"
	"site_builder_backend/pkg/redis"

	"github.com/gin-gonic/gin"
)

// Register registers all HTTP routes to the router
func Register(l *logger.ZapLogger, router *gin.Engine, pgClient *postgres.Postgres, redisClient *redis.Redis, esClient *elasticsearch.Elasticsearch) {
	// User routes
	userGroup := router.Group("User")
	user_router.UserRegister(userGroup, pgClient)

	addressGroup := router.Group("Address")
	user_router.AddressRegister(addressGroup, pgClient)
}
