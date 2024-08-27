package contexts

import (
	"context"

	"github.com/phongloihong/event-driven-mono/libs/configLoader"
	"github.com/phongloihong/event-driven-mono/libs/database/mongoLoader"
	"github.com/phongloihong/event-driven-mono/libs/log"
	"github.com/phongloihong/event-driven-mono/libs/models"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/config"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Log             log.ILogger
	Cfg             config.ConfigData
	MongoConn       *mongo.Database
	QueueConnection *amqp091.Connection
	CartRepo        mongoLoader.IRepository[models.CartModel]
}

func NewServiceContext(ctx context.Context) *ServiceContext {
	logger := log.NewLogger()

	cfg, err := configLoader.LoadConfig[config.ConfigData](".")
	if err != nil {
		logger.Panic(ctx, err)
	}

	mongoConn, err := getMongoConnection(cfg)
	if err != nil {
		logger.Panic(ctx, err)
	}

	return &ServiceContext{
		Log:       logger,
		Cfg:       cfg,
		MongoConn: mongoConn,
		CartRepo:  models.NewCartRepository(mongoConn, "carts"),
	}
}

func getMongoConnection(config config.ConfigData) (*mongo.Database, error) {
	mongoConfig := mongoLoader.MongoConfig{
		Host:     config.DatabaseHost,
		Port:     config.DatabasePort,
		Username: config.DatabaseUsername,
		Password: config.DatabasePassword,
		DBName:   config.DatabaseName,
		Ssl:      config.DataBaseSSL,
	}
	return mongoLoader.NewConnection(mongoConfig)
}
