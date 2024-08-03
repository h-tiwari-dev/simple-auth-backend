package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe user object.
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Username     string    `db:"username" json:"username" validate:"required,lte=255"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password_hash" validate:"required"`
	Active       bool      `db:"active" json:"active" validate:"required"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
