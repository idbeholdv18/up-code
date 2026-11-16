package events

import (
	"time"

	"github.com/google/uuid"
)

const (
	TopicUserRegistered = "user.registered"
)

type Envelope struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	OccuredAt time.Time `json:"occured_at"`
	Payload   any       `json:"payload"`
}

func NewEnvelope(eventType string, payload any) Envelope {
	return Envelope{
		EventID:   uuid.NewString(),
		EventType: eventType,
		OccuredAt: time.Now(),
		Payload:   payload,
	}
}

type UserRegisteredPayload struct {
	UserID           int64  `json:"user_id"`
	Email            string `json:"email"`
	FullName         string `json:"full_name"`
	ConfirmationLink string `json:"confirmation_link"`
}
