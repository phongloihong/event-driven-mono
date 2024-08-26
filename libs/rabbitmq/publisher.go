package rabbitmq

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/phongloihong/event-driven-mono/libs/log"
	"github.com/phongloihong/event-driven-mono/libs/otel"
	"github.com/phongloihong/event-driven-mono/libs/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type IPublisher interface {
	PublishMessage(msg interface{}) error
}

type publisher struct {
	cfg    *RabbitMQConfig
	conn   *amqp.Connection
	log    log.ILogger
	tracer trace.Tracer
	ctx    context.Context
}

func (p publisher) PublishMessage(msg interface{}) error {
	data, err := json.Marshal(msg)
	if err != nil {
		p.log.Errorf(p.ctx, "Failed to marshal message: %v", err)
		return err
	}

	typeName := reflect.TypeOf(msg).Elem().Name()
	snakeTypeName := utils.ToSnakeCase(typeName)

	ctx, span := p.tracer.Start(p.ctx, typeName)
	defer span.End()

	channel, err := p.conn.Channel()
	if err != nil {
		p.log.Error(ctx, "Failed to open a channel")
		return err
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(
		p.cfg.ExchangeName, // Name
		p.cfg.Kind,         // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		p.log.Error(ctx, "Failed to declare an exchange")
		return err
	}

	// Inject the context in the headers
	headers := otel.InjectAMQPHeaders(ctx)
	correlationId := otel.GetTraceID(ctx)
	publishingMsg := amqp.Publishing{
		Body:          data,
		ContentType:   "application/json",
		DeliveryMode:  amqp.Persistent, // Save message to disk
		MessageId:     uuid.New().String(),
		Timestamp:     time.Now(),
		CorrelationId: correlationId,
		Headers:       headers,
	}
	err = channel.PublishWithContext(
		ctx,
		p.cfg.ExchangeName, // exchange
		snakeTypeName,      // routing key
		false,              // mandatory
		false,              // immediate
		publishingMsg,      // message
	)
	if err != nil {
		p.log.Error(ctx, "Failed to publish a message")
		return err
	}

	h, err := json.Marshal(headers)
	if err != nil {
		p.log.Error(ctx, "Failed to marshal headers")
		return err
	}

	p.log.Infof(ctx, "Published message: %s", publishingMsg.Body)
	span.SetAttributes(attribute.Key("message-id").String(publishingMsg.MessageId))
	span.SetAttributes(attribute.Key("correlation-id").String(publishingMsg.CorrelationId))
	span.SetAttributes(attribute.Key("exchange").String(snakeTypeName))
	span.SetAttributes(attribute.Key("kind").String(p.cfg.Kind))
	span.SetAttributes(attribute.Key("content-type").String("application/json"))
	span.SetAttributes(attribute.Key("timestamp").String(publishingMsg.Timestamp.String()))
	span.SetAttributes(attribute.Key("body").String(string(publishingMsg.Body)))
	span.SetAttributes(attribute.Key("headers").String(string(h)))

	return nil
}

func NewPublisher(ctx context.Context, cfg *RabbitMQConfig, conn *amqp.Connection, log log.ILogger, tracer trace.Tracer) IPublisher {
	return &publisher{
		cfg:    cfg,
		conn:   conn,
		log:    log,
		tracer: tracer,
		ctx:    ctx,
	}
}
