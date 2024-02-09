package ticket

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/domain/category"
	"github.com/nafisalfiani/p3-final-project/product-service/domain/region"
	ticketDom "github.com/nafisalfiani/p3-final-project/product-service/domain/ticket"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

type Interface interface {
	List(ctx context.Context, filter entity.Ticket) ([]entity.Ticket, error)
	Get(ctx context.Context, filter entity.Ticket) (entity.Ticket, error)
	Create(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error)
	Update(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error)
	Delete(ctx context.Context, ticket entity.Ticket) error
}

type ticket struct {
	logger   log.Interface
	ticket   ticketDom.Interface
	category category.Interface
	region   region.Interface
}

func Init(logger log.Interface, prd ticketDom.Interface, cat category.Interface, reg region.Interface) Interface {
	return &ticket{
		logger:   logger,
		ticket:   prd,
		category: cat,
		region:   reg,
	}
}

func (t *ticket) List(ctx context.Context, filter entity.Ticket) ([]entity.Ticket, error) {
	tickets, err := t.ticket.List(ctx, filter)
	if err != nil {
		t.logger.Error(ctx, err)
		return nil, err
	}

	return tickets, nil
}

func (t *ticket) Get(ctx context.Context, filter entity.Ticket) (entity.Ticket, error) {
	ticket, err := t.ticket.Get(ctx, filter)
	if err != nil {
		t.logger.Error(ctx, err)
		return ticket, err
	}

	return ticket, nil
}

func (t *ticket) Create(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error) {
	cat, err := t.category.Get(ctx, ticket.Category)
	if err != nil {
		return entity.Ticket{}, err
	}

	reg, err := t.region.Get(ctx, ticket.Region)
	if err != nil {
		return entity.Ticket{}, err
	}

	ticket.Category = cat
	ticket.Region = reg

	newTicket, err := t.ticket.Create(ctx, ticket)
	if err != nil {
		t.logger.Error(ctx, err)
		return newTicket, err
	}

	return newTicket, nil
}

func (t *ticket) Update(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error) {
	newTicket, err := t.ticket.Update(ctx, ticket)
	if err != nil {
		t.logger.Error(ctx, err)
		return newTicket, err
	}

	return newTicket, nil
}

func (t *ticket) Delete(ctx context.Context, ticket entity.Ticket) error {
	if err := t.ticket.Delete(ctx, ticket); err != nil {
		t.logger.Error(ctx, err)
		return err
	}

	return nil
}
