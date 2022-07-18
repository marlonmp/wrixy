package redis_test

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func init() {

	uri := "redis://root@localhost:6379/0"

	opt, err := redis.ParseURL(uri)

	if err != nil {
		panic(err)
	}

	client = redis.NewClient(opt)

	err = client.Ping(context.Background()).Err()

	if err != nil {
		panic(err)
	}
}

func Client() *redis.Client {
	return client
}
