package port

import (
	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/user/domain"
)

type AccountService interface {
	Get(uuid.UUID) (domain.User, error)

	Update(uuid.UUID, domain.User) (domain.User, error)

	Delete(uuid.UUID) (domain.User, error)
}
