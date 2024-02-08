package category

import (
	"context"

	categoryDom "github.com/nafisalfiani/p3-final-project/product-service/domain/category"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

type Interface interface {
	List(ctx context.Context) ([]entity.Category, error)
	Get(ctx context.Context, filter entity.Category) (entity.Category, error)
	Create(ctx context.Context, category entity.Category) (entity.Category, error)
	Update(ctx context.Context, category entity.Category) (entity.Category, error)
	Delete(ctx context.Context, category entity.Category) error
}

type category struct {
	category categoryDom.Interface
}

func Init(prd categoryDom.Interface) Interface {
	return &category{
		category: prd,
	}
}

func (c *category) List(ctx context.Context) ([]entity.Category, error) {
	return c.category.List(ctx)
}

func (c *category) Get(ctx context.Context, filter entity.Category) (entity.Category, error) {
	return c.category.Get(ctx, filter)
}

func (c *category) Create(ctx context.Context, category entity.Category) (entity.Category, error) {
	return c.category.Create(ctx, category)
}

func (c *category) Update(ctx context.Context, category entity.Category) (entity.Category, error) {
	return c.category.Update(ctx, category)
}

func (c *category) Delete(ctx context.Context, category entity.Category) error {
	return c.category.Delete(ctx, category)
}
