package routing

import (
	"site_builder_backend/configs"
	"site_builder_backend/internal/infrastructures/impl/auth"
	"site_builder_backend/internal/infrastructures/impl/db/mysql/user_repo"
	"site_builder_backend/internal/interfaces/auth_inter"
	"site_builder_backend/internal/interfaces/db/repositories/user_repo_inter"
	"site_builder_backend/internal/presentation/middlewares"
	"site_builder_backend/pkg/elasticsearch"
	"site_builder_backend/pkg/logger"
	"site_builder_backend/pkg/postgres"
	"site_builder_backend/pkg/redis"
	"time"
)

const (
	_defaultConnectTimeout = 5 * time.Second
)

type Services struct {
	Logger           *logger.ZapLogger
	ElasticClient    *elasticsearch.Elasticsearch
	RedisClient      *redis.Redis
	PostgresClient   *postgres.Postgres
	AuthMiddleware   *middlewares.AuthMiddleware
	jwtService       auth_inter.JWTService
	UserReadRepo     user_repo_inter.UserReadRepository
	UserWriteRepo    user_repo_inter.UserWriteRepository
	AddressWriteRepo user_repo_inter.AddressWriteRepository
	AddressReadRepo  user_repo_inter.AddressReadRepository
}

func NewServiceRegistration(cfg *configs.Config, l *logger.ZapLogger) *Services {
	jwtService := auth.NewJWTService(cfg)
	var pgClient *postgres.Postgres
	var redisClient *redis.Redis
	var esClient *elasticsearch.Elasticsearch
	var err error

	// Initialize PostgreSQL with GORM
	pgClient, err = postgres.New(cfg.PG.URL, l, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal("app - Run - postgres.New: %v", err)
	}
	defer func() {
		if pgClient != nil {
			pgClient.Close()
		}
	}()

	// Initialize Redis with multiple databases
	redisClient, err = redis.New(cfg.Redis, l,
		redis.ConnectTimeout(_defaultConnectTimeout),
		redis.PoolSize(cfg.Redis.PoolSize),
		redis.MinIdleConns(cfg.Redis.MinIdleConns))
	if err != nil {
		l.Fatal("app - Run - redis.New: %v", err)
	}
	defer func() {
		if redisClient != nil {
			redisClient.Close()
		}
	}()

	// Initialize Elasticsearch
	esClient, err = elasticsearch.New(cfg.Elasticsearch, l,
		elasticsearch.ConnectTimeout(_defaultConnectTimeout),
		elasticsearch.SniffEnabled(cfg.Elasticsearch.SniffEnabled),
		elasticsearch.Retries(3))
	if err != nil {
		l.Fatal("app - Run - elasticsearch.New: %v", err)
	}
	// Log enabled Elasticsearch indexes
	l.Info("Elasticsearch initialized with indexes: %v", esClient.EnabledIndexes())

	userReadRepo := user_repo.NewUserReadRepository(pgClient.DB, l)
	userWriteRepo := user_repo.NewUserWriteRepository(pgClient.DB, l)

	addressReadRepo := user_repo.NewAddressReadRepository(pgClient.DB, l)
	addressWriteRepo := user_repo.NewAddressWriteRepository(pgClient.DB, l)

	return &Services{
		//System Injection
		Logger:         l,
		ElasticClient:  esClient,
		RedisClient:    redisClient,
		PostgresClient: pgClient,
		//Service Injection
		jwtService:     jwtService,
		AuthMiddleware: middlewares.NewAuthMiddleware(jwtService),
		//Repository injection
		UserReadRepo:     userReadRepo,
		UserWriteRepo:    userWriteRepo,
		AddressReadRepo:  addressReadRepo,
		AddressWriteRepo: addressWriteRepo,
	}
}
