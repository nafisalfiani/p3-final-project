package ticket

import (
	"context"

	ticketDom "github.com/nafisalfiani/p3-final-project/product-service/domain/ticket"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

type Interface interface {
	List(ctx context.Context) ([]entity.Ticket, error)
	Get(ctx context.Context, filter entity.Ticket) (entity.Ticket, error)
	Create(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error)
	Update(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error)
	Delete(ctx context.Context, ticket entity.Ticket) error
}

type ticket struct {
	ticket ticketDom.Interface
}

func Init(prd ticketDom.Interface) Interface {
	return &ticket{
		ticket: prd,
	}
}

func (t *ticket) List(ctx context.Context) ([]entity.Ticket, error) {
	return t.ticket.List(ctx)
}

func (t *ticket) Get(ctx context.Context, filter entity.Ticket) (entity.Ticket, error) {
	return t.ticket.Get(ctx, filter)
}

func (t *ticket) Create(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error) {
	return t.ticket.Create(ctx, ticket)
}

func (t *ticket) Update(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error) {
	return t.ticket.Update(ctx, ticket)
}

func (t *ticket) Delete(ctx context.Context, ticket entity.Ticket) error {
	return t.ticket.Delete(ctx, ticket)
}
