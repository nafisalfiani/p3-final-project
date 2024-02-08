package ticket

import (
	"context"
	"time"

	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

func (t *ticket) getCache(ctx context.Context, key string) (entity.Ticket, error) {
	var product entity.Ticket
	productStr, err := t.cache.Get(ctx, key)
	if err != nil {
		return product, err
	}

	if err := t.json.Unmarshal([]byte(productStr), &product); err != nil {
		return product, err
	}

	return product, nil
}

func (t *ticket) setCache(ctx context.Context, key string, product entity.Ticket) error {
	productJson, err := t.json.Marshal(product)
	if err != nil {
		return err
	}

	if err := t.cache.SetEX(ctx, key, string(productJson), time.Hour); err != nil {
		return err
	}

	return nil
}
