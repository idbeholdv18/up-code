package events

import "context"

type Publisher interface {
	Publish(ctx context.Context, routingKey string, event Envelope) error
}
