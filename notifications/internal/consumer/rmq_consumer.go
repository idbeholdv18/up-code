package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/idbeholdv18/up-code/notifications/internal/dispatcher"
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

	if len(msg.Data) == 0 {
		_ = delivery.Discard(ctx, nil)
		return
	}

	var envelope dispatcher.Envelope
	if err := json.Unmarshal(msg.Data[0], &envelope); err != nil {
		log.Printf("malformed envelope json (%v), discarding", err)
		_ = delivery.Discard(ctx, nil)
		return
	}

	if err := c.dispatcher.Dispatch(ctx, envelope); err != nil {
		if errors.Is(err, dispatcher.ErrInvalidPayload) {
			log.Printf("invalid payload for event id=%s (%v)", envelope.EventID, err)
			_ = delivery.Discard(ctx, nil)
			return
		}

		log.Printf("transient error processing event id=%s (%v)", envelope.EventID, err)
		time.Sleep(2 * time.Second)
		_ = delivery.Requeue(ctx)
		return
	}
	if err := delivery.Accept(ctx); err != nil {
		log.Printf("failed to accept event id=%s: (%v)", envelope.EventID, err)
	}
}
