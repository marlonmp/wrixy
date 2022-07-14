package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func init() {

	uri := os.Getenv("REDIS_URI")

	opt, err := redis.ParseURL(uri)

	if err != nil {
		panic(err)
	}

	client = redis.NewClient(opt)
}

func Client() *redis.Client {
	return client
}
