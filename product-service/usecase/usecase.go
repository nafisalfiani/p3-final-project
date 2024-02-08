package usecase

import (
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/domain"
	"github.com/nafisalfiani/p3-final-project/product-service/usecase/category"
	"github.com/nafisalfiani/p3-final-project/product-service/usecase/ticket"
)

type Usecases struct {
	Ticket   ticket.Interface
	Category category.Interface
}

func Init(logger log.Interface, dom *domain.Domains) *Usecases {
	return &Usecases{
		Ticket:   ticket.Init(dom.Ticket),
		Category: category.Init(dom.Category),
	}
}
