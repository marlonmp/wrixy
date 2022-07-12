package port

import (
	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/user/domain"
)

type UserFilter struct {
	UUID uuid.UUID `query:"uuid"`

	Username string `query:"username"`
	Email    string `query:"email"`

	HasDeletedAt bool `query:"has_deleted_at"`

	Offset uint `query:"o"`
	Limit  uint `query:"l"`
}

type UserRepo interface {
	Find(UserFilter) (*[]domain.User, error)

	FindOne(UserFilter) (domain.User, error)

	InsertOne(domain.User) (domain.User, error)

	UpdateOne(UserFilter, domain.User) (domain.User, error)

	DeleteOne(UserFilter) (domain.User, error)
}

type UserService interface {
	List(*UserFilter) (*[]domain.User, error)

	Get(UserFilter) (domain.User, error)

	Post(domain.User) (domain.User, error)

	Update(uuid.UUID, domain.User) (domain.User, error)

	Delete(uuid.UUID) (domain.User, error)
}
