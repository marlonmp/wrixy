package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID

	Nickname string

	Email    string
	Username string

	Password string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
