package config

import (
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/cache"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/nosql"
	"github.com/nafisalfiani/p3-final-project/lib/security"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc"
)

type Application struct {
	Auth     auth.Config     `env:"AUTH"`
	Log      log.Config      `env:"LOG"`
	Security security.Config `env:"SECURITY"`
	NoSql    nosql.Config    `env:"NO_SQL"`
	Cache    cache.Config    `env:"CACHE"`
	Broker   broker.Config   `env:"BROKER"`
	Grpc     grpc.Config     `env:"GRPC"`
}
