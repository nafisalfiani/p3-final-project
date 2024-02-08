package mailer

import (
	"context"

	mailerDom "github.com/nafisalfiani/p3-final-project/notification-service/domain/mailer"
	"github.com/nafisalfiani/p3-final-project/notification-service/entity"
)

type Interface interface {
	Send(ctx context.Context, content entity.Email) error
}

type mailer struct {
	mailer mailerDom.Interface
}

func Init(mailerDom mailerDom.Interface) Interface {
	return &mailer{
		mailer: mailerDom,
	}
}

func (e *mailer) Send(ctx context.Context, content entity.Email) error {
	return e.mailer.Send(ctx, content)
}
