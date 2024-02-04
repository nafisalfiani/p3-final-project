package role

import (
	"context"
	"time"

	"github.com/nafisalfiani/p3-final-project/account-service/entity"
)

func (r *role) getCache(ctx context.Context, key string) (entity.Role, error) {
	var role entity.Role
	roleStr, err := r.cache.Get(ctx, key)
	if err != nil {
		return role, err
	}

	if err := r.json.Unmarshal([]byte(roleStr), &role); err != nil {
		return role, err
	}

	return role, nil
}

func (r *role) setCache(ctx context.Context, key string, role entity.Role) error {
	roleJson, err := r.json.Marshal(role)
	if err != nil {
		return err
	}

	if err := r.cache.SetEX(ctx, key, string(roleJson), time.Hour); err != nil {
		return err
	}

	return nil
}
