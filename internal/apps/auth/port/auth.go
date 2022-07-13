package port

import (
	"time"

	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
)

type AuthService interface {
	SignIn(credentials domain.Credentials) (sid string, err error)

	Verify(sid string) (session domain.Session, err error)

	Refresh(sid string, duration time.Duration) (newSID string, err error)

	SignOut(sid string) error
}
