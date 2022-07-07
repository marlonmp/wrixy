package domain

import (
	"time"

	"github.com/google/uuid"
	bDomain "github.com/marlonmp/wrixy/internal/apps/book/domain"
	pDomain "github.com/marlonmp/wrixy/internal/apps/permission/domain"
	uDomain "github.com/marlonmp/wrixy/internal/apps/user/domain"
)

type BookList struct {
	UUID uuid.UUID

	OwnerUUID uuid.UUID
	Owner     *uDomain.User

	Title     string
	BooksUUID *[]uuid.UUID
	Books     *[]bDomain.Book

	Access pDomain.Access

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
