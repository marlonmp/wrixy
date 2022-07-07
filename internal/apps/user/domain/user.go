package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID

	Names string

	Email    string
	Username string

	Password string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
