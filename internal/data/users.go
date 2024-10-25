// Filename: internal/data/users.go

package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinezmoises/quiz3/internal/validator"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	Version   int32     `json:"version"`
}

type UserModel struct {
	DB *sql.DB
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Email != "", "email", "must be provided")
	v.Check(user.FullName != "", "full_name", "must be provided")
}

func (u UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users (email, full_name)
		VALUES ($1, $2)
		RETURNING id, created_at, version
	`
	args := []any{user.Email, user.FullName}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return u.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
}

func (u UserModel) Get(id int64) (*User, error) {
	if id < 1 {
		return nil, ErrUserNotFound
	}

	query := `
		SELECT id, created_at, email, full_name, version
		FROM users
		WHERE id = $1
	`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.CreatedAt, &user.Email, &user.FullName, &user.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u UserModel) Update(user *User) error {
	query := `
		UPDATE users
		SET email = $1, full_name = $2, version = version + 1
		WHERE id = $3
		RETURNING version
	`
	args := []any{user.Email, user.FullName, user.ID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return u.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
}

func (u UserModel) Delete(id int64) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := u.DB.ExecContext(ctx, query, id)
	return err
}
