package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/services"
)

func (store *Store) GetUser(username string) (*domain.User, error) {
	query := `
	SELECT id, created_at, name, email, password
	FROM users
	WHERE username = $1`
	var user *domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := store.db.QueryRowContext(ctx, query, username).Scan(&user.Id,
		&user.CreatedAt,
		&user.Username,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (store *Store) AddUser(user *domain.User) (*domain.User, error) {
	query := `
INSERT INTO users (id,username, email, password)
VALUES ($1, $2, $3,$4)
RETURNING id, created_at, version`
	args := []interface{}{user.Id, user.Username, user.Email, user.Password}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := store.db.QueryRowContext(ctx, query, args...).Scan(&user.CreatedAt)
	if err != nil {
		fmt.Print(err)
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return nil, services.ErrEmailUnique
		default:
			return nil, err
		}
	}
	return nil, nil
}

func (store *Store) DeleteUser(id uuid.UUID) error {
	query := `
		DELETE FROM users
		WHERE id = $1`

	result, err := store.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("err no record found it says") //ErrRecordNotFound
	}
	return nil
}
