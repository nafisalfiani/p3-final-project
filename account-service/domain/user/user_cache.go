package user

import (
	"context"
	"time"

	"github.com/nafisalfiani/p3-final-project/account-service/entity"
)

func (u *user) getCache(ctx context.Context, key string) (entity.User, error) {
	var user entity.User
	userStr, err := u.cache.Get(ctx, key)
	if err != nil {
		return user, err
	}

	if err := u.json.Unmarshal([]byte(userStr), &user); err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) setCache(ctx context.Context, key string, user entity.User) error {
	userJson, err := u.json.Marshal(user)
	if err != nil {
		return err
	}

	if err := u.cache.SetEX(ctx, key, string(userJson), time.Hour); err != nil {
		return err
	}

	return nil
}
