package region

import (
	"context"

	regionDom "github.com/nafisalfiani/p3-final-project/product-service/domain/region"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

type Interface interface {
	List(ctx context.Context) ([]entity.Region, error)
	Get(ctx context.Context, filter entity.Region) (entity.Region, error)
	Create(ctx context.Context, region entity.Region) (entity.Region, error)
	Update(ctx context.Context, region entity.Region) (entity.Region, error)
	Delete(ctx context.Context, region entity.Region) error
}

type region struct {
	region regionDom.Interface
}

func Init(reg regionDom.Interface) Interface {
	return &region{
		region: reg,
	}
}

func (r *region) List(ctx context.Context) ([]entity.Region, error) {
	return r.region.List(ctx)
}

func (r *region) Get(ctx context.Context, filter entity.Region) (entity.Region, error) {
	return r.region.Get(ctx, filter)
}

func (r *region) Create(ctx context.Context, region entity.Region) (entity.Region, error) {
	return r.region.Create(ctx, region)
}

func (r *region) Update(ctx context.Context, region entity.Region) (entity.Region, error) {
	return r.region.Update(ctx, region)
}

func (r *region) Delete(ctx context.Context, region entity.Region) error {
	return r.region.Delete(ctx, region)
}
