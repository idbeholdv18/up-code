package mail

type ConfirmationEmailEvent struct {
	EventID          string `json:"event_id"`
	Email            string `json:"email"`
	FullName         string `json:"full_name"`
	ConfirmationLink string `json:"confirmation_link"`
}
