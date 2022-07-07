package port

import (
	"github.com/google/uuid"
	"github.com/marlonmp/wrixy/internal/apps/booklist/domain"
)

type BookListFilter struct {
	UUID uuid.UUID `query:"uuid"`

	Title         string   `query:"title"`
	BookTitle     string   `query:"book_title"`
	CategoriesStr []string `query:"categories"`

	AuthorUUID uuid.UUID `query:"author"`

	HasDeletedAt bool `query:"has_deleted_at"`
}

type BookListRepo interface {
	Find(*BookListFilter) (*[]domain.BookList, error)

	FindOne(*BookListFilter) (domain.BookList, error)

	InsertOne(domain.BookList) (domain.BookList, error)

	UpdateOne(uuid.UUID, domain.BookList) (domain.BookList, error)

	DeleteOne(uuid.UUID) (domain.BookList, error)
}

type BookListService interface {
	List(*BookListFilter) (*[]domain.BookList, error)

	Get(*BookListFilter) (domain.BookList, error)

	Post(domain.BookList) (domain.BookList, error)

	Update(uuid.UUID, domain.BookList) (domain.BookList, error)

	Delete(uuid.UUID) (domain.BookList, error)
}
