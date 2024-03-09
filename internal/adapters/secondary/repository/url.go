package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/services"
)

func (store *Store) GetUrls(userId uuid.UUID) ([]*domain.Url, error) {

	query := `
SELECT id, user_id,short_url,original_url,created_at
FROM urls
WHERE user_id = $1
ORDER BY created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := store.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	urls := []*domain.Url{}

	for rows.Next() {

		var url domain.Url

		err := rows.Scan(
			&url.Id,
			&url.UserId,
			&url.ShortUrl,
			&url.OriginalUrl,
			&url.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		urls = append(urls, &url)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil

}
func (store *Store) GetUrl(shortUrl string) (*domain.Url, error) {
	query := `
SELECT id, user_id,short_url,original_url,created_at
FROM urls
WHERE id = $1`

	var url domain.Url

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
			return nil, services.ErrUrlNotFound
		default:
			return nil, err
		}
	}

	return &url, nil
}
func (store *Store) AddUrl(url *domain.Url) (*domain.Url, error) {

	query := `
	INSERT INTO urls (id, user_id,short_url,original_url)
	VALUES ($1, $2, $3,$4)
	RETURNING created_at`
	args := []interface{}{url.Id, url.UserId, url.ShortUrl, url.OriginalUrl}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := store.db.QueryRowContext(ctx, query, args...).Scan(&url.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "urls_short_url_key"`:
			return nil, services.ErrDepulicateShortUrl
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
		return services.ErrUrlNotFound
	}
	return nil
}
