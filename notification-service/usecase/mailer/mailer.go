package mailer

import (
	"context"
	"fmt"

	"github.com/nafisalfiani/p3-final-project/lib/header"
	mailerDom "github.com/nafisalfiani/p3-final-project/notification-service/domain/mailer"
	"github.com/nafisalfiani/p3-final-project/notification-service/entity"
)

type Config struct {
	ApiGatewayUrl string
}

type Interface interface {
	SendRegistrationEmail(ctx context.Context, user entity.User) error
	SendTransactionEmail(ctx context.Context) error
}

type mailer struct {
	mailer mailerDom.Interface
	config Config
}

func Init(config Config, mailerDom mailerDom.Interface) Interface {
	return &mailer{
		mailer: mailerDom,
		config: config,
	}
}

func (m *mailer) SendRegistrationEmail(ctx context.Context, user entity.User) error {
	content := entity.Email{
		Body:        fmt.Sprintf(`Verify email here: %v/auth/v1/verify-email/%v`, m.config.ApiGatewayUrl, user.Id),
		BodyType:    header.ContentTypePlain,
		Subject:     "Ketson - Email Verification",
		SenderName:  "Ketson",
		SenderEmail: "no-reply@ketson.com",
		Recipients: entity.Recipient{
			ToEmails:  []string{user.Email},
			BCCEmails: []string{"nafisa.alfiani.ica@gmail.com"},
		},
	}

	return m.mailer.Send(ctx, content)
}

func (m *mailer) SendTransactionEmail(ctx context.Context) error {
	// TODO: compose email
	content := entity.Email{}

	return m.mailer.Send(ctx, content)
}
