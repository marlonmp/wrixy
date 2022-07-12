package port

import "github.com/google/uuid"

// JWT blacklist
type JWTRepo interface {
	Add(uuid.UUID) error
	Has(uuid.UUID) (bool, error)
}
