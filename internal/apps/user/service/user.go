package service

import (
	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/user/domain"
	"github.com/marlonmp/wrixy/internal/apps/user/port"
)

type userService struct {
	repo port.UserRepo
}

func UserService(repo port.UserRepo) port.UserService {
	return userService{repo}
}

func (us userService) List(filter *port.UserFilter) (*[]domain.User, error) {

	if filter == nil {
		filter = &port.UserFilter{}
	}

	filter.HasDeletedAt = false
	filter.Offset = 0
	filter.Limit = 10

	return us.repo.Find(*filter)
}

func (us userService) Get(filter port.UserFilter) (domain.User, error) {

	filter.HasDeletedAt = false

	return us.repo.FindOne(filter)
}

func (us userService) Post(user domain.User) (domain.User, error) {

	return us.repo.InsertOne(user)
}

func (us userService) Update(id uuid.UUID, user domain.User) (domain.User, error) {

	filter := port.UserFilter{
		UUID:         id,
		HasDeletedAt: false,
	}

	return us.repo.UpdateOne(filter, user)
}

func (us userService) Delete(id uuid.UUID) (domain.User, error) {

	filter := port.UserFilter{
		UUID:         id,
		HasDeletedAt: false,
	}

	return us.repo.DeleteOne(filter)
}
