package mail

import "context"

type Publisher interface {
	PublishConfirmation(ctx context.Context, event ConfirmationEmailEvent) error
}