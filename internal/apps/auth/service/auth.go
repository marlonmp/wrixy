package service

import (
	"time"

	authDomain "github.com/marlonmp/wrixy/internal/apps/auth/domain"
	authPort "github.com/marlonmp/wrixy/internal/apps/auth/port"
	userDomain "github.com/marlonmp/wrixy/internal/apps/user/domain"
	userPort "github.com/marlonmp/wrixy/internal/apps/user/port"
)

const defaultDuration time.Duration = time.Hour * 2

type authService struct {
	users userPort.UserRepo
}

func AuthService(repo userPort.UserRepo) authPort.AuthService {
	return authService{repo}
}

func (as authService) Authenticate(credentials authDomain.Credentials) (user userDomain.User, err error) {

	userFilter := userPort.UserFilter{
		Username: credentials.Identification,
	}

	user, err = as.users.FindOne(userFilter)

	if err != nil {
		return
	}

	if !matchPassword(user.Password, credentials.Password) {
		return userDomain.User{}, nil
	}

	return
}
