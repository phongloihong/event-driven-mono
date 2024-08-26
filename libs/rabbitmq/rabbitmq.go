package rabbitmq

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type RabbitMQConfig struct {
	Uri          string
	Password     string
	ExchangeName string
	Kind         string
}

func NewRabbitMQ(ctx context.Context, cfg *RabbitMQConfig) (*amqp.Connection, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 10 * time.Second
	maxRetries := 5

	var conn *amqp.Connection
	var err error

	err = backoff.Retry(func() error {
		conn, err = amqp.Dial(cfg.Uri)
		if err != nil {
			logrus.Errorf("Failed to connect to RabbitMQ: %v", err)
			return err
		}

		return nil
	}, backoff.WithMaxRetries(bo, uint64(maxRetries-1)))

	logrus.Info("Connected to RabbitMQ")

	go func() {
		<-ctx.Done()
		err := conn.Close()
		if err != nil {
			logrus.Error("Failed to close RabbitMQ connection")
		}

		logrus.Info("Closed RabbitMQ connection")
	}()

	return conn, err
}
