package port

import (
	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
	userDomain "github.com/marlonmp/wrixy/internal/apps/user/domain"
)

type AuthService interface {
	Authenticate(credentials domain.Credentials) (userDomain.User, error)
}
