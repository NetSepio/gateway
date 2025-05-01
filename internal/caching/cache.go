package caching

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"netsepio-gateway-v1.1/utils/load"
)

var (
	Rdb *redis.Client
	ctx = context.Background() // context.Background() is a function that returns a new context
)

func InitRedis() {
	address := load.Cfg.REDIS_HOST
	password := load.Cfg.REDIS_PASSWORD

	Rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // No password set
		DB:       0,        // Use default DB
		Protocol: 2,        // Connection protocol
	})
	if Rdb.Ping(ctx).Err() != nil {
		logrus.Fatal(Rdb.Ping(ctx).Err())
	}
}
