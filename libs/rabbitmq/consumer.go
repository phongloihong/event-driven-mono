package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/phongloihong/event-driven-mono/libs/log"
	"github.com/phongloihong/event-driven-mono/libs/otel"
	"github.com/phongloihong/event-driven-mono/libs/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type IConsumer[T any] interface {
	ConsumeMessage(msg interface{}, dependencies T) error
}

type Consumer[T any] struct {
	cfg     *RabbitMQConfig
	conn    *amqp.Connection
	log     log.ILogger
	handler func(queue string, msg amqp.Delivery, dependencies T) error
	tracer  trace.Tracer
	ctx     context.Context
}

func (c Consumer[T]) ConsumeMessage(msg interface{}, dependencies T) error {
	strName := strings.Split(runtime.FuncForPC(reflect.ValueOf(c.handler).Pointer()).Name(), ".")
	var consumerHandlerName = strName[len(strName)-1]

	ch, err := c.conn.Channel()
	if err != nil {
		c.log.Errorf(c.ctx, "Failed to open a channel %v", err)
		return err
	}

	typeName := reflect.TypeOf(msg).Name()
	snakeTypeName := utils.ToSnakeCase(typeName)

	err = ch.ExchangeDeclare(
		c.cfg.ExchangeName, // exchange name
		c.cfg.Kind,         // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		c.log.Errorf(c.ctx, "Failed to declare an exchange %v", err)
		return err
	}

	q, err := ch.QueueDeclare(
		fmt.Sprintf("%s_%s", snakeTypeName, "queue"), // name
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		c.log.Errorf(c.ctx, "Failed to declare a queue %v", err)
		return err
	}

	err = ch.QueueBind(
		q.Name, // queue name
		snakeTypeName,
		c.cfg.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		c.log.Errorf(c.ctx, "Failed to bind a queue %v", err)
		return err
	}

	deliveries, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		c.log.Error(c.ctx, "Failed to consume a message")
		return err
	}

	go func() {
		select {
		case <-c.ctx.Done():
			defer func(ch *amqp.Channel) {
				err := ch.Close()
				if err != nil {
					c.log.Errorf(c.ctx, "Failed to close a channel %v", err)
				}
			}(ch)

			c.log.Infof(c.ctx, "Consumer %s is stopped", consumerHandlerName)
		case delivery, ok := <-deliveries:
			if !ok {
				c.log.Error(c.ctx, "Failed to consume a message")
				return
			}

			// extract headers
			c.ctx = otel.ExtractAMQPHeaders(c.ctx, delivery.Headers)
			err := c.handler(q.Name, delivery, dependencies)
			if err != nil {
				c.log.Errorf(c.ctx, "Failed to handle a message %v", err)
			}

			_, span := c.tracer.Start(c.ctx, consumerHandlerName)
			h, err := json.Marshal(delivery.Headers)
			if err != nil {
				c.log.Errorf(c.ctx, "Failed to marshal headers %v", err)
				delivery.Nack(false, false)
				return
			}

			span.SetAttributes(attribute.Key("message-id").String(delivery.MessageId))
			span.SetAttributes(attribute.Key("correlation-id").String(delivery.CorrelationId))
			span.SetAttributes(attribute.Key("queue").String(q.Name))
			span.SetAttributes(attribute.Key("exchange").String(delivery.Exchange))
			span.SetAttributes(attribute.Key("routing-key").String(delivery.RoutingKey))
			span.SetAttributes(attribute.Key("ack").Bool(true))
			span.SetAttributes(attribute.Key("timestamp").String(delivery.Timestamp.String()))
			span.SetAttributes(attribute.Key("body").String(string(delivery.Body)))
			span.SetAttributes(attribute.Key("headers").String(string(h)))

			// Can not defer inside a for loop
			time.Sleep(1 * time.Microsecond)
			span.End()

			err = delivery.Ack(false)
			if err != nil {
				c.log.Errorf(c.ctx, "Failed to ack a message %v", err)
			}
		}
	}()

	c.log.Infof(c.ctx, "Waiting for messages in queue :%s. To exit press CTRL+C", q.Name)

	return nil
}

func NewConsumer[T any](ctx context.Context, cfg *RabbitMQConfig, conn *amqp.Connection, log log.ILogger, tracer trace.Tracer, handler func(queue string, msg amqp.Delivery, dependencies T) error) IConsumer[T] {
	return &Consumer[T]{ctx: ctx, cfg: cfg, conn: conn, log: log, tracer: tracer, handler: handler}
}
