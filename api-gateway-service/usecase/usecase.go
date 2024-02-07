package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/domain"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/security"
)

type Usecases struct {
}

func Init(logger log.Interface, sec security.Interface, validator *validator.Validate, dom *domain.Domains) *Usecases {
	return &Usecases{}
}
