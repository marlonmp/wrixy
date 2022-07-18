package service

import (
	"time"

	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
	"github.com/marlonmp/wrixy/internal/apps/auth/port"
)

type sessionService struct {
	sessions port.SessionRepo
}

func SessionService(repo port.SessionRepo) port.SessionService {
	return sessionService{repo}
}

func (ss sessionService) Set(session domain.Session, expiration time.Duration) (sid string, err error) {

	sid = RandomSID()

	err = ss.sessions.Set(sid, session, expiration)

	return
}

func (ss sessionService) Get(sid string) (domain.Session, error) {

	return ss.sessions.Get(sid)
}

func (ss sessionService) ExtendDuration(sid string, expiration time.Duration) error {

	return ss.sessions.Expire(sid, expiration)
}

func (ss sessionService) Remove(sid string) error {

	return ss.sessions.Del(sid)
}
