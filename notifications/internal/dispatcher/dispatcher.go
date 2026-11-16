package dispatcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/idbeholdv18/up-code/notifications/internal/channel"
)

var ErrInvalidPayload = errors.New("ivalid event payload")

type Envelope struct {
	EventID   string          `json:"event_id"`
	EventType string          `json:"event_type"`
	OccuredAt time.Time       `json:"occured_at"`
	Payload   json.RawMessage `json:"payload"`
}

type UserRegisteredPayload struct {
	UserID           int64  `json:"user_id"`
	Email            string `json:"email"`
	FullName         string `json:"full_name"`
	ConfirmationLink string `json:"confirmation_link"`
}

type Dispatcher struct {
	senders map[string]channel.Sender
}

func NewDispatcher(senders ...channel.Sender) *Dispatcher {
	m := make(map[string]channel.Sender, len(senders))
	for _, s := range senders {
		m[s.Name()] = s
	}
	return &Dispatcher{
		senders: m,
	}
}

func (d *Dispatcher) Dispatch(ctx context.Context, envelope Envelope) error {
	switch envelope.EventType {
	case "user.registered":
		return d.handleUserRegistered(ctx, envelope.Payload)
	default:
		log.Printf("skipping unhandled event type: %s (event id: %s)", envelope.EventType, envelope.EventID)
		return nil
	}
}

func (d *Dispatcher) handleUserRegistered(ctx context.Context, rawPayload json.RawMessage) error {
	var p UserRegisteredPayload
	if err := json.Unmarshal(rawPayload, &p); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidPayload, err)
	}

	n := channel.Notification{
		Recipient: p.Email,
		Subject:   "LMS registration confirmation",
		Body: fmt.Sprintf(
			"<h3>Welcome, %s</h3><p>To finish registration tap to link:</p><a href=\"%s\">Confirm</a>",
			p.FullName,
			p.ConfirmationLink,
		),
	}

	emailSender, ok := d.senders["email"]
	if !ok {
		return fmt.Errorf("Sender email is not configured")
	}

	return emailSender.Send(ctx, n)
}
