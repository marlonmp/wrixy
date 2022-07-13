package port

import (
	"time"

	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
)

type SessionRepo interface {
	Set(session domain.Session, duration time.Duration) (sid string, err error)
	Get(sid string) (domain.Session, error)
	Del(sid string) error
}
