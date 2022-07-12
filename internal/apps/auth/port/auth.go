package port

import (
	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
)

type AuthService interface {
	SignIn(c domain.Credentials) (token string, err error)

	Verify(token string) error

	SignOut(tokenID uuid.UUID) error
}
