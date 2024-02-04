package usecase

import (
	"github.com/nafisalfiani/p3-final-project/account-service/domain"
	"github.com/nafisalfiani/p3-final-project/account-service/usecase/role"
	"github.com/nafisalfiani/p3-final-project/account-service/usecase/user"
	"github.com/nafisalfiani/p3-final-project/lib/log"
)

type Usecases struct {
	User user.Interface
	Role role.Interface
}

func Init(logger log.Interface, dom *domain.Domains) *Usecases {
	return &Usecases{
		User: user.Init(logger, dom.User),
		Role: role.Init(logger, dom.Role),
	}
}
