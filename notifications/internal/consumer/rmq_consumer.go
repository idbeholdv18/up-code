package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"time"

	"github.com/idbeholdv18/up-code/notifications/internal/dispatcher"
	"github.com/idbeholdv18/up-code/shared/httpkit"
	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

type RMQConsumer struct {
	consumer   *rmq.Consumer
	dispatcher *dispatcher.Dispatcher
}

func NewRMQConsumer(consumer *rmq.Consumer, dispatcher *dispatcher.Dispatcher) *RMQConsumer {
	return &RMQConsumer{
		consumer:   consumer,
		dispatcher: dispatcher,
	}
}

func (c *RMQConsumer) Run(ctx context.Context) {
	log.Println("Notification consumer started...")

	for {
		delivery, err := c.consumer.Receive(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				log.Println("Notification consumer stopped")
				return
			}
			log.Printf("recieve error: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		c.process(ctx, delivery)
	}
}

func (c *RMQConsumer) process(ctx context.Context, delivery rmq.IDeliveryContext) {
	msg := delivery.Message()

	var reqID string
	if msg.ApplicationProperties != nil {
		if val, ok := msg.ApplicationProperties["x-request-id"].(string); ok {
			reqID = val
		}
	}

	if reqID == "" {
		reqID = httpkit.GenerateRequestID()
	}

	ctx = httpkit.WithRequestID(ctx, reqID)

	slog.InfoContext(ctx, "recieved notification event")

	if len(msg.Data) == 0 {
		_ = delivery.Discard(ctx, nil)
		return
	}

	var envelope dispatcher.Envelope
	if err := json.Unmarshal(msg.Data[0], &envelope); err != nil {
		slog.ErrorContext(ctx, "failed to unmarshal message",
			slog.String("err_msg", err.Error()),
		)
		_ = delivery.Discard(ctx, nil)
		return
	}

	if err := c.dispatcher.Dispatch(ctx, envelope); err != nil {
		if errors.Is(err, dispatcher.ErrInvalidPayload) {
			slog.ErrorContext(ctx, "invalid payload",
				slog.String("event_id", envelope.EventID),
				slog.String("err_msg", err.Error()),
			)
			_ = delivery.Discard(ctx, nil)
			return
		}

		slog.ErrorContext(ctx, "transient error",
			slog.String("event_id", envelope.EventID),
			slog.String("err_msg", err.Error()),
		)
		time.Sleep(2 * time.Second)
		_ = delivery.Requeue(ctx)
		return
	}

	if err := delivery.Accept(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to accept",
			slog.String("event_id", envelope.EventID),
			slog.String("err_msg", err.Error()),
		)
	}
}
