package events

import (
	"context"
	"encoding/json"
	"fmt"

	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

type rmqPublisher struct {
	publisher    *rmq.Publisher
	exchangeName string
}

func NewRmqPublisher(publisher *rmq.Publisher) Publisher {
	return &rmqPublisher{
		publisher: publisher,
	}
}

func (p *rmqPublisher) Publish(ctx context.Context, routingKey string, event Envelope) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal event %s: %v", event.EventType, err)
	}

	msg, err := rmq.NewMessageWithAddress(payload, &rmq.ExchangeAddress{
		Exchange: p.exchangeName,
		Key:      routingKey,
	})

	if err != nil {
		return fmt.Errorf("create message with addr for %s: %v", event.EventType, err)
	}

	res, err := p.publisher.Publish(ctx, msg)
	if err != nil {
		return fmt.Errorf("publish event %s: %v", event.EventType, err)
	}

	switch res.Outcome.(type) {
	case *rmq.StateAccepted:
		return nil
	default:
		return fmt.Errorf("unexpected pubish outcome for event %s: %v", event.EventType, err)
	}
}
