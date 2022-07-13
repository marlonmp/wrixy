package service

import (
	"time"

	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
	"github.com/marlonmp/wrixy/internal/apps/auth/port"
	uport "github.com/marlonmp/wrixy/internal/apps/user/port"
)

const defaultDuration time.Duration = time.Hour * 2

type authService struct {
	repo  port.SessionRepo
	urepo uport.UserRepo
}

func AuthService(r port.SessionRepo, ur uport.UserRepo) port.AuthService {
	return authService{r, ur}
}

func (as authService) SignIn(credentials domain.Credentials) (string, error) {

	ufilter := uport.UserFilter{
		Username: credentials.Identification,
	}

	u, err := as.urepo.FindOne(ufilter)

	if err != nil {
		return "", err
	}

	session := domain.Session{Identifier: u.Username}

	return as.repo.Set(session, defaultDuration)
}

func (as authService) SignOut(sid string) error {

	return as.repo.Del(sid)
}

func (as authService) Verify(sid string) (domain.Session, error) {

	return as.repo.Get(sid)
}

func (as authService) Refresh(sid string, duration time.Duration) (newSID string, err error) {

	session, err := as.repo.Get(sid)

	if err != nil {
		return "", err
	}

	return as.repo.Set(session, duration)
}
