package mailer

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/lib/email"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/notification-service/entity"
)

type Interface interface {
	Send(ctx context.Context, content entity.Email) error
}

type mailer struct {
	logger log.Interface
	email  email.Interface
}

func Init(log log.Interface, email email.Interface) Interface {
	return &mailer{
		logger: log,
		email:  email,
	}
}

func (m *mailer) Send(ctx context.Context, content entity.Email) error {
	params := email.SendEmailParams{
		Body:        content.Body,
		BodyType:    content.BodyType,
		Subject:     content.Subject,
		SenderName:  content.SenderName,
		SenderEmail: content.SenderEmail,
		Recipients: email.Recipient{
			ToEmails:  content.Recipients.ToEmails,
			CCEmails:  content.Recipients.CCEmails,
			BCCEmails: content.Recipients.BCCEmails,
		},
		Attachments: content.Attachments,
		Headers:     content.Headers,
	}

	return m.email.SendEmail(ctx, params)
}
