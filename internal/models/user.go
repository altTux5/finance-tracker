package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/altTux5/expense-tracker/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     *string // nullable
	Phone     *string // nullable
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(user User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, email, phone, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := db.DB.Exec(query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func GetUserByEmailOrPhone(email, phone string) (*User, error) {
	var user User
	var row *sql.Row

	if email != "" {
		query := `
			SELECT id, first_name, last_name, email, phone, password, created_at, updated_at
			FROM users
			WHERE  email = $1
			LIMIT 1
		`
		row = db.DB.QueryRow(query, email)
	} else if phone != "" {
		query := `
			SELECT id, first_name, last_name, email, phone, password, created_at, updated_at
			FROM users
			WHERE  phone = $1
			LIMIT 1
		`
		row = db.DB.QueryRow(query, phone)
	}

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
