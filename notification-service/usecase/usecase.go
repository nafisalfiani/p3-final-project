package usecase

import (
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/notification-service/domain"
	"github.com/nafisalfiani/p3-final-project/notification-service/usecase/mailer"
)

type Usecases struct {
	Mailer mailer.Interface
}

func Init(apiGwUrl string, logger log.Interface, dom *domain.Domains) *Usecases {
	return &Usecases{
		Mailer: mailer.Init(mailer.Config{
			ApiGatewayUrl: apiGwUrl,
		}, dom.Mailer),
	}
}
