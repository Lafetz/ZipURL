package ports

import (
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

type UrlRepository interface {
	GetUrls(userId uuid.UUID) ([]*domain.Url, error)
	GetUrl(shortUrl string) (*domain.Url, error)
	AddUrl(*domain.Url) (*domain.Url, error)
	UpdateUrl(*domain.Url) error
	DeleteUrl(uuid.UUID) error
}
