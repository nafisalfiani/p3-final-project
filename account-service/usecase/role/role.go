package role

import (
	"context"

	roleDom "github.com/nafisalfiani/p3-final-project/account-service/domain/role"
	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/log"
)

type role struct {
	logger log.Interface
	role   roleDom.Interface
}

type Interface interface {
	List(ctx context.Context) ([]entity.Role, error)
	Get(ctx context.Context, role entity.Role) (entity.Role, error)
	Create(ctx context.Context, role entity.Role) (entity.Role, error)
}

func Init(logger log.Interface, roleDom roleDom.Interface) Interface {
	return &role{
		logger: logger,
		role:   roleDom,
	}
}

func (r *role) List(ctx context.Context) ([]entity.Role, error) {
	return r.role.List(ctx)
}

func (r *role) Get(ctx context.Context, role entity.Role) (entity.Role, error) {
	return r.role.Get(ctx, role)
}

func (r *role) Create(ctx context.Context, role entity.Role) (entity.Role, error) {
	return r.role.Create(ctx, role)
}
