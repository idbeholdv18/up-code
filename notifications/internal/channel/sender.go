package channel

import "context"

type Notification struct {
	Recipient string
	Subject   string
	Body      string
}

type Sender interface {
	Name() string
	Send(ctx context.Context, n Notification) error
}
