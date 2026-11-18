package store

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s UserStore) Create(ctx context.Context, user *User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = ""

	query := `
	INSERT INTO users(username, password, email)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`

	if err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		string(hashPassword),
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}
