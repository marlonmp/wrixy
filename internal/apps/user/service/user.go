// TODO: add user and uuid validations and its tests
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

func (us userService) List(filter port.UserFilter) ([]domain.User, error) {

	if filter.Limit < 1 {
		filter.Limit = 10
		filter.Offset = 0
	}

	if filter.Offset < 0 {
		filter.Offset = 0
	}

	filter.IsDeleted = false

	return us.repo.Find(filter)
}

func (us userService) Get(filter port.UserFilter) (domain.User, error) {

	filter.IsDeleted = false

	return us.repo.FindOne(filter)
}

func (us userService) Post(user domain.User) (domain.User, error) {

	return us.repo.InsertOne(user)
}

func (us userService) Update(id uuid.UUID, user domain.User) (domain.User, error) {

	return us.repo.UpdateOne(id, user)
}

func (us userService) Delete(id uuid.UUID) (domain.User, error) {

	return us.repo.DeleteOne(id)
}
