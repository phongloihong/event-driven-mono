package contexts

import (
	"context"

	"github.com/phongloihong/event-driven-mono/libs/configLoader"
	"github.com/phongloihong/event-driven-mono/libs/database/mongoLoader"
	"github.com/phongloihong/event-driven-mono/libs/log"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/config"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Log             log.ILogger
	Cfg             config.ConfigData
	MongoConn       *mongo.Database
	QueueConnection *amqp091.Connection
}

func NewServiceContext(ctx context.Context) ServiceContext {
	logger := log.NewLogger()

	config, err := configLoader.LoadConfig[config.ConfigData](".")
	if err != nil {
		logger.Panic(ctx, err)
	}

	mongoConfig := mongoLoader.MongoConfig{
		Host:     config.DatabaseHost,
		Port:     config.DatabasePort,
		Username: config.DatabaseUsername,
		Password: config.DatabasePassword,
		DBName:   config.DatabaseName,
		Ssl:      config.DataBaseSSL,
	}
	mongoConn, err := mongoLoader.NewConnection(mongoConfig)
	if err != nil {
		logger.Panic(ctx, err)
	}

	return ServiceContext{
		Log:       logger,
		Cfg:       config,
		MongoConn: mongoConn,
		// QueueConnection: queueConn,
	}
}
