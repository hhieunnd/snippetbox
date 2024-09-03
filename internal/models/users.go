package models

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *pgxpool.Pool
}

// Insert new record into table users.
func (u *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, password, created)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP AT TIME ZONE 'UTC')`

	_, err = u.DB.Exec(context.Background(), stmt, name, email, string(hashedPassword))

	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == "23505" && strings.Contains(pgError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}

		return err
	}

	return nil
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Check if uses exists with a specific ID.
func (u *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
