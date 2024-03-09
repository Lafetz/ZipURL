package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

func (store *Store) GetUrls(userId uuid.UUID) ([]*domain.Url, error) {

	query := `
SELECT id, user_id,short_url,original_url,created_at, 
FROM movies
ORDER BY created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := store.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	urls := []*domain.Url{}

	for rows.Next() {

		var url *domain.Url

		err := rows.Scan(
			url.Id,
			url.UserId,
			url.ShortUrl,
			url.OriginalUrl,
			url.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil

}
func (store *Store) GetUrl(shortUrl string) (*domain.Url, error) {
	query := `
SELECT id, user_id,short_url,original_url,created_at, 
FROM movies
WHERE id = $1`

	var url *domain.Url

	err := store.db.QueryRow(query, shortUrl).Scan(
		url.Id,
		url.UserId,
		url.ShortUrl,
		url.OriginalUrl,
		url.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("err no record found it says")
		default:
			return nil, err
		}
	}

	return url, nil
}
func (store *Store) AddUrl(url *domain.Url) (*domain.Url, error) {

	query := `
	INSERT INTO users (id, user_id,short_url,original_url,created_at)
	VALUES ($1, $2, $3,$4,$5)
	RETURNING id, created_at, version`
	args := []interface{}{url.Id, url.UserId, url.ShortUrl, url.OriginalUrl}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := store.db.QueryRowContext(ctx, query, args...).Scan(&url.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return nil, err
		default:
			return nil, err
		}
	}
	return nil, nil
}

func (store *Store) DeleteUrl(shorturl string) error {
	query := `
	DELETE FROM urls
	WHERE short_url = $1`

	result, err := store.db.Exec(query, shorturl)
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
