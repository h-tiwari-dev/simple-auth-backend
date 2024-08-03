package queries

import (
	"sample-auth-backend/app/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQuery struct {
	*sqlx.DB
}

// GetUsers method for getting all users.
func (q *UserQuery) GetUsers() ([]models.User, error) {
	users := []models.User{}
	query := "SELECT * FROM users;"

	err := q.Select(&users, query)
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetUser method for getting one user by given ID.
func (q *UserQuery) GetUser(id uuid.UUID) (models.User, error) {
	user := models.User{}
	query := "SELECT * FROM users WHERE id = $1;"

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser method for creating a user by given User object.
func (q *UserQuery) CreateUser(u *models.User) error {
	query := "INSERT INTO users (id, username, password_hash, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := q.Exec(query, u.ID, u.Username, u.PasswordHash, u.Email, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser method for updating a user by given User object.
func (q *UserQuery) UpdateUser(id uuid.UUID, u *models.User) error {
	query := "UPDATE users SET username = $2, email = $3, updated_at = $4 WHERE id = $1;"

	_, err := q.Exec(query, id, u.Username, u.Email, u.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser method for deleting a user by given ID.
func (q *UserQuery) DeleteUser(id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1;"

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQuery) GetUserWithEmailOrUserName(email string, username string) (*models.User, error) {
	user := models.User{}

	err := q.QueryRowx("SELECT * FROM users us WHERE us.username=$1 OR us.email=$2", username, email).StructScan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
