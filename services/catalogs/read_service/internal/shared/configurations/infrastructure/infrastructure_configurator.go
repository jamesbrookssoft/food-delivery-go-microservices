package infrastructure

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core/serializer"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core/serializer/json"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/grpc"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	messageBus "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/bus"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/producer"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing"
	rabbitmqConfigurations "github.com/mehdihadeli/store-golang-microservice-sample/pkg/rabbitmq/configurations"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/read_service/config"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/read_service/internal/shared/web/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
)

type InfrastructureConfigurations struct {
	Log                          logger.Logger
	Cfg                          *config.Config
	Validator                    *validator.Validate
	Producer                     producer.Producer
	PgConn                       *pgxpool.Pool
	Gorm                         *gorm.DB
	Metrics                      *CatalogsServiceMetrics
	TraceProvider                *trace.TracerProvider
	Esdb                         *esdb.Client
	MongoClient                  *mongo.Client
	GrpcClient                   grpc.GrpcClient
	ElasticClient                *elasticsearch.Client
	Redis                        redis.UniversalClient
	MiddlewareManager            cutomMiddlewares.CustomMiddlewares
	EventSerializer              serializer.EventSerializer
	RabbitMQConfigurationBuilder rabbitmqConfigurations.RabbitMQConfigurationBuilder
	RabbitMQBus                  messageBus.Bus
}

type InfrastructureConfigurator interface {
	ConfigInfrastructures(ctx context.Context) (*InfrastructureConfigurations, error, func())
}

type infrastructureConfigurator struct {
	log logger.Logger
	cfg *config.Config
}

func NewInfrastructureConfigurator(log logger.Logger, cfg *config.Config) InfrastructureConfigurator {
	return &infrastructureConfigurator{log: log, cfg: cfg}
}

func (ic *infrastructureConfigurator) ConfigInfrastructures(ctx context.Context) (*InfrastructureConfigurations, error, func()) {
	infrastructure := &InfrastructureConfigurations{Cfg: ic.cfg, Log: ic.log, Validator: validator.New()}

	metrics := ic.configCatalogsMetrics()
	infrastructure.Metrics = metrics

	cleanup := []func(){}

	grpcClient, err := grpc.NewGrpcClient(ic.cfg.GRPC)
	if err != nil {
		return nil, err, nil
	}
	cleanup = append(cleanup, func() {
		_ = grpcClient.Close()
	})
	infrastructure.GrpcClient = grpcClient

	traceProvider, err := tracing.AddOtelTracing(ic.cfg.OTel)
	if err != nil {
		return nil, err, nil
	}
	cleanup = append(cleanup, func() {
		_ = traceProvider.Shutdown(ctx)
	})
	infrastructure.TraceProvider = traceProvider

	mongoClient, err, mongoCleanup := ic.configMongo(ctx)
	if err != nil {
		return nil, err, nil
	}
	cleanup = append(cleanup, mongoCleanup)
	infrastructure.MongoClient = mongoClient

	redis, err, redisCleanup := ic.configRedis(ctx)
	if err != nil {
		return nil, err, nil
	}
	cleanup = append(cleanup, redisCleanup)
	infrastructure.Redis = redis

	infrastructure.EventSerializer = json.NewJsonEventSerializer()

	infrastructure.RabbitMQConfigurationBuilder = rabbitmqConfigurations.NewRabbitMQConfigurationBuilder()

	return infrastructure, nil, func() {
		for _, c := range cleanup {
			defer c()
		}
	}
}
