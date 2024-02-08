package domain

import (
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/email"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"github.com/nafisalfiani/p3-final-project/notification-service/domain/mailer"
)

type Domains struct {
	Mailer mailer.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, broker broker.Interface, mail email.Interface) *Domains {
	return &Domains{
		Mailer: mailer.Init(logger, mail),
	}
}
