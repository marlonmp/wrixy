package port

import (
	"time"

	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
)

type SessionRepo interface {
	Set(key string, session domain.Session, expiration time.Duration) error

	Get(key string) (domain.Session, error)

	Expire(key string, expiration time.Duration) error

	Del(key string) error
}

type SessionService interface {
	Set(session domain.Session, expiration time.Duration) (sid string, err error)

	Get(sid string) (domain.Session, error)

	ExtendDuration(sid string, expiration time.Duration) error

	Remove(sid string) error
}
