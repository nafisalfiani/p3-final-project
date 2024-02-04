package nosql

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	MaxIdleConn uint64        `env:"NO_SQL_MAX_IDLE_CONN"`
	MaxIdleTime time.Duration `env:"NO_SQL_MAX_IDLE_TIME"`
	DSN         string        `env:"NO_SQL_DSN"`
}

func Init(cfg Config) (*mongo.Client, error) {
	opts := options.Client().
		ApplyURI(cfg.DSN).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)).
		SetMaxPoolSize(cfg.MaxIdleConn).
		SetMinPoolSize(1).
		SetMaxConnIdleTime(cfg.MaxIdleTime)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	ctx, cleanup := context.WithTimeout(context.Background(), cfg.MaxIdleTime)
	defer cleanup()

	if err := client.Ping(ctx, readpref.Nearest()); err != nil {
		return nil, err
	}

	return client, nil
}
