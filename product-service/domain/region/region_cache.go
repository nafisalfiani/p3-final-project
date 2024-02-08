package region

import (
	"context"
	"time"

	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

func (r *region) getCache(ctx context.Context, key string) (entity.Region, error) {
	var category entity.Region
	categoryStr, err := r.cache.Get(ctx, key)
	if err != nil {
		return category, err
	}

	if err := r.json.Unmarshal([]byte(categoryStr), &category); err != nil {
		return category, err
	}

	return category, nil
}

func (r *region) setCache(ctx context.Context, key string, category entity.Region) error {
	categoryJson, err := r.json.Marshal(category)
	if err != nil {
		return err
	}

	if err := r.cache.SetEX(ctx, key, string(categoryJson), time.Hour); err != nil {
		return err
	}

	return nil
}
