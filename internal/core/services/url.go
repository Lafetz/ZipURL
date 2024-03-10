package services

import (
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/ports"
)

type UrlServiceApi interface {
	GetUrls(uuid.UUID) ([]*domain.Url, error)
	GetUrl(string) (*domain.Url, error)
	AddUrl(*domain.Url) (*domain.Url, error)
	DeleteUrl(string, uuid.UUID) error
}
type UrlService struct {
	repo ports.UrlRepository
}

func NewUrlService(repo ports.UrlRepository) *UrlService {
	return &UrlService{
		repo: repo,
	}
}
func (srv *UrlService) GetUrls(userId uuid.UUID) ([]*domain.Url, error) {
	return srv.repo.GetUrls(userId)
}
func (srv *UrlService) GetUrl(shortUrl string) (*domain.Url, error) {
	return srv.repo.GetUrl(shortUrl)
}

func (srv *UrlService) AddUrl(url *domain.Url) (*domain.Url, error) {

	for i := 0; i < 10; i++ { // retry upto 10 times if same shorturl id is found
		id := uuid.New().String()
		truncatedID := id[:7]
		url.ShortUrl = truncatedID
		_, err := srv.repo.AddUrl(url)
		if err != nil {
			switch {
			case err == ErrDepulicateShortUrl:
				continue
			default:
				return nil, err
			}
		}
		return url, nil

	}
	return nil, ErrUrlDepulicateRetry
}
func (srv *UrlService) DeleteUrl(shortUrl string, userId uuid.UUID) error {

	return srv.repo.DeleteUrl(shortUrl, userId)
}
