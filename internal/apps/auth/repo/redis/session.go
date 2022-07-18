// TODO: create redis errors for not found
package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
	"github.com/marlonmp/wrixy/internal/apps/auth/port"
)

type sessionRepo struct {
	client *redis.Client
}

func SessionRepo(repo *redis.Client) port.SessionRepo {
	return sessionRepo{repo}
}

func (sr sessionRepo) Set(key string, sess domain.Session, exp time.Duration) error {

	ctx := context.Background()

	conn := sr.client.Conn(ctx)

	defer conn.Close()

	sessBuf := marsharl(sess)

	status := conn.SetEX(ctx, key, sessBuf, exp)

	err := status.Err()

	return err
}

func (sr sessionRepo) Get(key string) (sess domain.Session, err error) {

	ctx := context.Background()

	conn := sr.client.Conn(ctx)

	defer conn.Close()

	val, err := conn.Get(ctx, key).Bytes()

	if err != nil {
		return
	}

	return unmarshal(val)
}

func (sr sessionRepo) Expire(key string, exp time.Duration) error {

	ctx := context.Background()

	conn := sr.client.Conn(ctx)

	defer conn.Close()

	ok, err := conn.Expire(ctx, key, exp).Result()

	if !ok {
		return errors.New("not found")
	}

	return err
}

func (sr sessionRepo) Del(key string) error {

	ctx := context.Background()

	conn := sr.client.Conn(ctx)

	defer conn.Close()

	n, err := conn.Del(ctx, key).Result()

	if n < 1 {
		return errors.New("not found")
	}

	return err
}

func marsharl(sess domain.Session) []byte {
	buf, _ := jsoniter.Marshal(sess)

	return buf
}

func unmarshal(data []byte) (sess domain.Session, err error) {
	err = jsoniter.Unmarshal(data, &sess)

	return
}
