package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
	"github.com/marlonmp/wrixy/internal/apps/auth/port"
	"github.com/marlonmp/wrixy/internal/apps/auth/repo/redis"
	"github.com/marlonmp/wrixy/internal/apps/auth/service"
)

var (
	sessionRepo port.SessionRepo

	key  = service.RandomSID()
	sess = domain.Session{Identifier: "manolo268"}
	exp  = 2 * time.Second
)

func init() {
	sessionRepo = redis.SessionRepo(client)
}

func TestSet(t *testing.T) {

	err := sessionRepo.Set(key, sess, exp)

	if err != nil {
		t.Fatal("session set err:", err)
	}
}

func TestGet(t *testing.T) {

	sessionRepo.Set(key, sess, exp)

	obtainedSession, err := sessionRepo.Get(key)

	if err != nil {
		t.Fatal("session set err:", err)
	}

	if sess.Identifier != obtainedSession.Identifier {
		t.Error("invalid identifier")
	}
}

func TestExpire(t *testing.T) {
	sessionRepo.Set(key, sess, exp)

	ctx := context.Background()

	conn := client.Conn(ctx)

	defer conn.Close()

	err := sessionRepo.Expire(key, exp*2)

	if err != nil {
		t.Fatal("session expire err:", err)
	}

	newExp := conn.TTL(ctx, key).Val()

	if newExp < exp {
		t.Error("invalid expiration time")
	}
}

func TestDel(t *testing.T) {
	sessionRepo.Set(key, sess, exp)

	ctx := context.Background()

	conn := client.Conn(ctx)

	defer conn.Close()

	err := sessionRepo.Del(key)

	if err != nil {
		t.Fatal("session del err:", err)
	}
}
