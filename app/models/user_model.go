package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe user object.
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required"`
	Name         string    `db:"name" json:"name" validate:"lte=255"`
	Username     string    `db:"username" json:"username" validate:"lte=255"`
	Email        string    `db:"email" json:"email" validate:"omitempty,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password" validate:""`
	PhoneNumber  string    `db:"phone_number" json:"phone_number" validate:""`
	Active       bool      `db:"active" json:"active" validate:""`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	LoginType    string    `db:"login_type" json:"login_type" validate:"required,oneof=simple google github facebook"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToUserResponse converts User to UserResponse.
func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type SignIn struct {
	Username string `json:"username" validate:"lte=255"`
	Email    string `json:"email" validate:"email,lte=255"`
	Password string `json:"password" validate:"required"`
}
