package service

import (
	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/auth/domain"
	"github.com/marlonmp/wrixy/internal/apps/auth/port"
	uport "github.com/marlonmp/wrixy/internal/apps/user/port"
)

type authService struct {
	repo  port.JWTRepo
	urepo uport.UserRepo
}

func AuthService(r port.JWTRepo, ur uport.UserRepo) port.AuthService {
	return authService{r, ur}
}

func (as authService) SignIn(credentials domain.Credentials) (token string, err error) {

	ufilter := uport.UserFilter{
		Username: credentials.Identification,
	}

	u, err := as.urepo.FindOne(ufilter)

	if err != nil {
		return
	}

	// TODO: apply jwt generate token

	// TODO: use `u: domain.User`

	return
}

func (as authService) Verify(token string) error {

	// TODO: apply jwt verify token

	hasID, err := as.repo.Has(id)

	if err != nil {
		// TODO: return fiber.StatusInternalServerError: 500
		return nil
	}

	if hasID {
		// TODO: return fiber.StatusForbidden: 403
		return nil
	}

	return nil
}

func (as authService) SignOut(id uuid.UUID) error {

	return as.repo.Add(id)
}
