package port

import "github.com/marlonmp/wrixy/internal/apps/category/domain"

type CategoryFilter struct {
	Name string `query:"name"`
}

type CategoryRepo interface {
	Find(*CategoryFilter) (*[]domain.Category, error)
}

type CategoryService interface {
	List(*CategoryFilter) (*[]domain.Category, error)
}
