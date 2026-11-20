package channel

import (
	"context"
	"fmt"
	"log/slog"
	"net/smtp"
)

type SMTPSender struct {
	addr string
	from string
}

func NewSMTPSender(addr string, from string) Sender {
	return &SMTPSender{
		addr: addr,
		from: from,
	}
}

func (s *SMTPSender) Name() string {
	return "email"
}

func (s *SMTPSender) Send(ctx context.Context, n Notification) error {
	subject := fmt.Sprintf("Subject: %s\r\n", n.Subject)
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	msg := []byte(subject + mime + n.Body)

	if err := smtp.SendMail(s.addr, nil, s.from, []string{n.Recipient}, msg); err != nil {
		return fmt.Errorf("smtp send mail to %s: %v", n.Recipient, err)
	}

	slog.InfoContext(ctx, "confirmation email has been sent")

	return nil
}
