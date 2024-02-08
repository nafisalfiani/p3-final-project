package category

import (
	"context"
	"time"

	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

func (c *category) getCache(ctx context.Context, key string) (entity.Category, error) {
	var category entity.Category
	categoryStr, err := c.cache.Get(ctx, key)
	if err != nil {
		return category, err
	}

	if err := c.json.Unmarshal([]byte(categoryStr), &category); err != nil {
		return category, err
	}

	return category, nil
}

func (c *category) setCache(ctx context.Context, key string, category entity.Category) error {
	categoryJson, err := c.json.Marshal(category)
	if err != nil {
		return err
	}

	if err := c.cache.SetEX(ctx, key, string(categoryJson), time.Hour); err != nil {
		return err
	}

	return nil
}
