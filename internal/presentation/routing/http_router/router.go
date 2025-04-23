package http_router

import (
	"site_builder_backend/internal/presentation/routing"
	"site_builder_backend/pkg/elasticsearch"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/postgres"
	"site_builder_backend/pkg/redis"

	"github.com/gin-gonic/gin"
)

func Register(l *logger.ZapLogger, router *gin.Engine, pgClient *postgres.Postgres, redisClient *redis.Redis, esClient *elasticsearch.Elasticsearch) {
	services := routing.NewServices(router, pgClient, redisClient, esClient, l)

	userRouter := NewUserRouter(services)
	userRouter.UserRegister()
	userRouter.AddressRegister()

	blogRouter := NewBlogRouter(services)
	blogRouter.ArticleRegister()
	blogRouter.CategoryRegister()

}
