package port

import (
	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/book/domain"
)

type BookFilter struct {
	UUID uuid.UUID `query:"uuid"`

	Title         string   `query:"title"`
	Abstract      string   `query:"abstract"`
	CategoriesStr []string `query:"categories"`

	AuthorUUID uuid.UUID `query:"author"`

	HasDeletedAt bool `query:"has_deleted_at"`
}

type BookRepo interface {
	Find(*BookFilter) (*[]domain.Book, error)

	FindOne(*BookFilter) (domain.Book, error)

	InsertOne(domain.Book) (domain.Book, error)

	UpdateOne(uuid.UUID, domain.Book) (domain.Book, error)

	DeleteOne(uuid.UUID) (domain.Book, error)
}

type BookService interface {
	List(*BookFilter) (*[]domain.Book, error)

	Get(*BookFilter) (domain.Book, error)

	Post(domain.Book) (domain.Book, error)

	Update(uuid.UUID, domain.Book) (domain.Book, error)

	Delete(uuid.UUID) (domain.Book, error)
}
