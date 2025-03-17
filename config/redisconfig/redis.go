package redisconfig

import (
	"context"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	Rdb *redis.Client
	ctx = context.Background() // context.Background() is a function that returns a new context
)

func InitRedis() {
	address := envconfig.EnvVars.REDIS_HOST
	password := envconfig.EnvVars.REDIS_PASSWORD

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
