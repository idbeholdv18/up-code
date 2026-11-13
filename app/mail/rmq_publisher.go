package mail

import (
	"context"
	"encoding/json"
	"fmt"

	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

type rmqPublisher struct {
	Publisher *rmq.Publisher
}

func NewRmqPublisher(publisher *rmq.Publisher) Publisher {
	return &rmqPublisher{
		Publisher: publisher,
	}
}

func (p *rmqPublisher) PublishConfirmation(ctx context.Context, event ConfirmationEmailEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal confirmation event: %v", err)
	}

	msg := rmq.NewMessage(payload)

	res, err := p.Publisher.Publish(ctx, msg)
	if err != nil {
		return fmt.Errorf("publish confirmation error: %v", err)
	}

	switch res.Outcome.(type) {
	case *rmq.StateAccepted:
		return nil
	default:
		return fmt.Errorf("unexpected publish outcome: %v", err)
	}
}
