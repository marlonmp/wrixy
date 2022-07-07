package domain

import (
	"time"

	"github.com/google/uuid"
	cDomain "github.com/marlonmp/wrixy/internal/apps/category/domain"
	pDomain "github.com/marlonmp/wrixy/internal/apps/permission/domain"
	uDomain "github.com/marlonmp/wrixy/internal/apps/user/domain"
)

type Book struct {
	UUID uuid.UUID

	AuthorUUID uuid.UUID
	Author     uDomain.User

	Title         string
	Abstract      string
	Pages         uint16
	CategoriesStr []string
	Categories    []cDomain.Category

	Access pDomain.Access

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
